package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/uschmann/clonybox/handler"
	"github.com/uschmann/clonybox/repos"
	"github.com/uschmann/clonybox/services"
	"github.com/uschmann/clonybox/storage"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	db, err := storage.OpenDb("test.db")
	melody := melody.New()
	settings := services.NewSettings(db)
	rfidChannel := make(chan string)
	broadcastService := services.NewBroadcastService(melody)
	playbackConfigRepo := repos.NewPlaybackConfigRepo(db)

	if err != nil {
		log.Fatal(err)
		return
	}

	env := &handler.Env{
		Db:                 db,
		SpotifyService:     services.NewSpotifyService("http://localhost:8080/callback"),
		Settings:           settings,
		PlaybackConfigRepo: playbackConfigRepo,
		RfidChannel:        rfidChannel,
		RfidObserver:       services.NewRfidObserver(rfidChannel, broadcastService, playbackConfigRepo, settings),
		BroadcastService:   broadcastService,
	}

	rfidReader := services.NewRfidReader()
	go rfidReader.StartReading(env.RfidChannel)
	go env.RfidObserver.Observe()

	r := gin.Default()

	handler.RegisterDeviceRoutes(r, env)
	handler.RegisterAuthRoutes(r, env)
	handler.RegisterPlaybackConfigRoutes(r, env)
	handler.RegisterSpotifyHandler(r, env)

	r.GET("/ws", func(c *gin.Context) {
		melody.HandleRequest(c.Writer, c.Request)
	})

	/*
		go func() {
			for rfid := range env.RfidChannel {
				env.BroadcastService.Broadcast("rfid.scanned", &map[string]interface{}{
					"rfid": rfid,
				})
			}
		}()
	*/

	if !env.Settings.Has("spotify.token") {
		env.SpotifyService.StartAuth()
	} else {
		tokenJson, _ := env.Settings.Get("spotify.token")
		var token oauth2.Token
		json.Unmarshal([]byte(tokenJson), &token)
		env.SpotifyService.Client = spotify.New(env.SpotifyService.Auth.Client(context.Background(), &token))
	}

	r.Run()
}
