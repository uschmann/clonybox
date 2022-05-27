package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
)

func RegisterAuthRoutes(r *gin.Engine, env *Env) {

	r.GET("/callback", func(c *gin.Context) {
		r := c.Request
		w := c.Writer

		tok, err := env.SpotifyService.Auth.Token(r.Context(), env.SpotifyService.State, r)
		if err != nil {
			http.Error(w, "Couldn't get token", http.StatusForbidden)
			log.Fatal(err)
		}
		if st := r.FormValue("state"); st != env.SpotifyService.State {
			http.NotFound(w, r)
			log.Fatalf("State mismatch: %s != %s\n", st, env.SpotifyService.State)
		}
		// use the token to get an authenticated client
		tokenJson, err := json.Marshal(tok)
		if err != nil {
			log.Fatal(err)
		}
		env.Settings.Set("spotify.token", string(tokenJson))
		client := spotify.New(env.SpotifyService.Auth.Client(r.Context(), tok))
		env.SpotifyService.ClientChannel <- client

		c.String(200, "Login complete!!!!!")
	})

}
