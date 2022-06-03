package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/uschmann/clonybox/handler"
	"github.com/uschmann/clonybox/repos"
	"github.com/uschmann/clonybox/services"
	rfidreader "github.com/uschmann/clonybox/services/rfidReader"
	"github.com/uschmann/clonybox/storage"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
	"gopkg.in/olahol/melody.v1"
)

//go:embed clonybox-frontend/dist/*
var res embed.FS

func serveStaticFiles(r *gin.Engine) {
	// Serve static embeded files
	r.GET("/", func(c *gin.Context) {
		c.FileFromFS("/clonybox-frontend/dist/index.htm", http.FS(res))
	})
	js, _ := fs.Sub(res, "clonybox-frontend/dist/js")
	r.StaticFS("/js", http.FS(js))
	css, _ := fs.Sub(res, "clonybox-frontend/dist/css")
	r.StaticFS("/css", http.FS(css))
}

func getCallbackUrl() string {
	port := os.Getenv("PORT")

	if port == "" || port == "80" {
		return "http://" + os.Getenv("HOST") + "/callback"
	}

	return "http://" + os.Getenv("HOST") + ":" + port + "/callback"
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db, err := storage.OpenDb("test.db")
	melody := melody.New()
	settings := services.NewSettings(db)
	rfidChannel := make(chan string)
	broadcastService := services.NewBroadcastService(melody)
	playbackConfigRepo := repos.NewPlaybackConfigRepo(db)

	callbackUrl := getCallbackUrl()
	spotifyService := services.NewSpotifyService(callbackUrl, settings)

	if err != nil {
		log.Fatal(err)
		return
	}

	env := &handler.Env{
		Db:                 db,
		SpotifyService:     spotifyService,
		Settings:           settings,
		PlaybackConfigRepo: playbackConfigRepo,
		RfidChannel:        rfidChannel,
		RfidObserver:       services.NewRfidObserver(rfidChannel, broadcastService, playbackConfigRepo, settings, spotifyService),
		BroadcastService:   broadcastService,
	}

	var rfidReader rfidreader.RfidReader = createRfIdReader(os.Getenv("RFID_TYPE")) //"/dev/input/event20"
	go rfidReader.StartReading(env.RfidChannel)
	go env.RfidObserver.Observe()

	r := gin.Default()

	serveStaticFiles(r)

	handler.RegisterDeviceRoutes(r, env)
	handler.RegisterAuthRoutes(r, env)
	handler.RegisterPlaybackConfigRoutes(r, env)
	handler.RegisterSpotifyHandler(r, env)
	handler.RegisterInfoHandler(r, env)

	r.GET("/api/ws", func(c *gin.Context) {
		melody.HandleRequest(c.Writer, c.Request)
	})

	// Authenticate spotify client

	if !env.Settings.Has("spotify.token") {
		env.SpotifyService.StartAuth()
	} else {
		tokenJson, _ := env.Settings.Get("spotify.token")
		var token oauth2.Token
		json.Unmarshal([]byte(tokenJson), &token)
		env.SpotifyService.Client = spotify.New(env.SpotifyService.Auth.Client(context.Background(), &token))

		fmt.Println("Token expires: ", token.Expiry)
	}

	go func() {
		for {
			time.Sleep(20 * time.Minute)
			if env.SpotifyService.Client != nil {
				env.SpotifyService.GetUser()
				token, err := env.SpotifyService.Client.Token()

				fmt.Println(token.Expiry)

				if err != nil {
					fmt.Println(err)
				}

				tokenJson, err := json.Marshal(token)
				if err != nil {
					fmt.Println(err)
				}

				env.Settings.Set("spotify.token", string(tokenJson))
			}
		}
	}()

	// Start server
	r.Run()
}

func createRfIdReader(readerType string) rfidreader.RfidReader {
	switch readerType {
	case "evdev":
		return rfidreader.NewEvdevRfIdReader(os.Getenv("RFID_EVDEV_FILE"))
	}

	return rfidreader.NewStdioRfidReader()
}
