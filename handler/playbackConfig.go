package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/uschmann/clonybox/models"
	"github.com/zmb3/spotify/v2"
	"gorm.io/datatypes"
)

func RegisterPlaybackConfigRoutes(r *gin.Engine, env *Env) {
	group := r.Group("/api/playback-config")
	{
		group.GET("/", env.indexPlaybackConfigs)
		group.POST("/", env.storePlaybackConfig)
	}
}

func (env *Env) indexPlaybackConfigs(c *gin.Context) {
	playbackConfigs, err := env.PlaybackConfigRepo.GetAllPlaybackConfigs()

	if err != nil {
		log.Fatal(err)
		c.AbortWithError(500, err)
	}

	c.JSON(200, playbackConfigs)
}

type storePlaybackConfigBody struct {
	Type      string `json:"type"`
	SpotifyId string `json:"spotify_id"`
}

func (env *Env) storePlaybackConfig(c *gin.Context) {
	var body storePlaybackConfigBody
	err := c.BindJSON(&body)

	if err != nil {
		fmt.Println(err)
		return
	}

	switch body.Type {
	case "album":
		album, _ := env.SpotifyService.Client.GetAlbum(context.Background(), spotify.ID(body.SpotifyId))
		metadata, _ := json.Marshal(album)
		playbackConfig := &models.PlaybackConfig{
			Name:       album.Name,
			Type:       "album",
			SpotifyUrl: string(album.URI),
			SpotifyId:  album.ID.String(),
			MetaData:   datatypes.JSON([]byte(metadata)),
		}

		env.PlaybackConfigRepo.CreatePlaybackConfig(playbackConfig)
		c.JSON(200, playbackConfig)
	case "playlist":
		playlist, _ := env.SpotifyService.Client.GetPlaylist(context.Background(), spotify.ID(body.SpotifyId))
		metadata, _ := json.Marshal(playlist)
		playbackConfig := &models.PlaybackConfig{
			Name:       playlist.Name,
			Type:       "album",
			SpotifyUrl: string(playlist.URI),
			SpotifyId:  playlist.ID.String(),
			MetaData:   datatypes.JSON([]byte(metadata)),
		}

		env.PlaybackConfigRepo.CreatePlaybackConfig(playbackConfig)
	case "track":
		track, _ := env.SpotifyService.Client.GetTrack(context.Background(), spotify.ID(body.SpotifyId))
		metadata, _ := json.Marshal(track)
		playbackConfig := &models.PlaybackConfig{
			Name:       track.Name,
			Type:       "album",
			SpotifyUrl: string(track.URI),
			SpotifyId:  track.ID.String(),
			MetaData:   datatypes.JSON([]byte(metadata)),
		}

		env.PlaybackConfigRepo.CreatePlaybackConfig(playbackConfig)
	}

	return

}
