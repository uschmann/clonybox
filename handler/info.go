package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
)

func RegisterInfoHandler(r *gin.Engine, env *Env) {
	info := r.Group("/api/info")
	{
		info.GET("/", env.indexInfo)
	}
}

func (env *Env) indexInfo(c *gin.Context) {
	authenticated := env.Settings.Has("spotify.token")
	var user *spotify.PrivateUser = nil
	var devices []spotify.PlayerDevice

	if authenticated {
		u, err := env.SpotifyService.GetUser()
		user = u

		if err != nil {
			fmt.Println(err)
		}

		d, err := env.SpotifyService.GetDevices()
		devices = d

		if err != nil {
			fmt.Println(err)
		}
	}

	c.JSON(200, gin.H{
		"user":             user,
		"is_authenticated": authenticated,
		"devices":          devices,
		"auth_url":         env.SpotifyService.GetAuthUrl(),
	})
}
