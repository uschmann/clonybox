package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

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
		group.GET("/:id", env.showPlaybackConfig)
		group.PUT("/:id", env.updatePlaybackConfig)
	}
}

func (env *Env) showPlaybackConfig(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	playbackConfig := env.PlaybackConfigRepo.GetPlaybackConfigById(id)

	c.JSON(200, playbackConfig)
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

type updatePlaybackConfigBody struct {
	SpotifyId string `json:"spotify_id"`
	Type      string `json:"type"`
}

func (env *Env) updatePlaybackConfig(c *gin.Context) {
	var body updatePlaybackConfigBody
	err := c.BindJSON(&body)

	if err != nil {
		fmt.Println(err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	playbackConfig := env.PlaybackConfigRepo.GetPlaybackConfigById(id)

	switch body.Type {
	case "album":
		album, _ := env.SpotifyService.Client.GetAlbum(context.Background(), spotify.ID(body.SpotifyId))
		metadata, _ := json.Marshal(album)
		playbackConfig.Type = "album"
		playbackConfig.Name = album.Name
		playbackConfig.SpotifyUrl = string(album.URI)
		playbackConfig.SpotifyId = album.ID.String()
		playbackConfig.MetaData = datatypes.JSON([]byte(metadata))

		env.PlaybackConfigRepo.UpdatePlaybackConfig(playbackConfig)
	case "playlist":
		playlist, _ := env.SpotifyService.Client.GetPlaylist(context.Background(), spotify.ID(body.SpotifyId))
		metadata, _ := json.Marshal(playlist)
		playbackConfig.Type = "playlist"
		playbackConfig.Name = playlist.Name
		playbackConfig.SpotifyUrl = string(playlist.URI)
		playbackConfig.SpotifyId = playlist.ID.String()
		playbackConfig.MetaData = datatypes.JSON([]byte(metadata))

		env.PlaybackConfigRepo.UpdatePlaybackConfig(playbackConfig)

	case "track":
		track, _ := env.SpotifyService.Client.GetTrack(context.Background(), spotify.ID(body.SpotifyId))
		metadata, _ := json.Marshal(track)
		playbackConfig.Type = "track"
		playbackConfig.Name = track.Name
		playbackConfig.SpotifyUrl = string(track.URI)
		playbackConfig.SpotifyId = track.ID.String()
		playbackConfig.MetaData = datatypes.JSON([]byte(metadata))

		env.PlaybackConfigRepo.UpdatePlaybackConfig(playbackConfig)
	}

	c.JSON(200, playbackConfig)
}
