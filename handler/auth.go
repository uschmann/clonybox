package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uschmann/clonybox/services"
	"github.com/zmb3/spotify/v2"
)

func RegisterAuthRoutes(r *gin.Engine) {

	r.GET("/callback", func(c *gin.Context) {
		r := c.Request
		w := c.Writer

		tok, err := services.SpotifyService.Auth.Token(r.Context(), services.SpotifyService.State, r)
		if err != nil {
			http.Error(w, "Couldn't get token", http.StatusForbidden)
			log.Fatal(err)
		}
		if st := r.FormValue("state"); st != services.SpotifyService.State {
			http.NotFound(w, r)
			log.Fatalf("State mismatch: %s != %s\n", st, services.SpotifyService.State)
		}
		// use the token to get an authenticated client
		client := spotify.New(services.SpotifyService.Auth.Client(r.Context(), tok))
		services.SpotifyService.ClientChannel <- client

		c.String(200, "Login complete!!!!!")
	})

}
