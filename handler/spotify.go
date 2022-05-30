package handler

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
)

func RegisterSpotifyHandler(r *gin.Engine, env *Env) {
	group := r.Group("/api/spotify")
	{
		group.POST("/search", env.searchSpotify)
	}
}

type searchSpotifyBody struct {
	Query string `json:"query"`
}

func (env *Env) searchSpotify(c *gin.Context) {
	var body searchSpotifyBody
	c.BindJSON(&body)

	result, err := env.SpotifyService.Client.Search(context.Background(), body.Query, spotify.SearchTypeAlbum|spotify.SearchTypePlaylist|spotify.SearchTypeTrack)

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, result)
}
