package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/uschmann/clonybox/handler"
	"github.com/uschmann/clonybox/services"
	"github.com/uschmann/clonybox/storage"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
)

func main() {

	db, err := storage.OpenDb("test.db")

	if err != nil {
		log.Fatal(err)
		return
	}

	env := &handler.Env{
		Db:             db,
		SpotifyService: services.NewSpotifyService("http://localhost:8080/callback"),
		Settings:       services.NewSettings(db),
	}

	r := gin.Default()

	handler.RegisterDeviceRoutes(r, env)
	handler.RegisterAuthRoutes(r, env)

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
