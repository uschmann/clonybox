package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
)

func RegisterDeviceRoutes(r *gin.Engine, env *Env) {

	device := r.Group("/api/device")
	{
		device.GET("/", env.deviceIndex)
		device.GET("/default", env.deviceGetDefault)
		device.POST("/default/:id", env.deviceSetDefault)
	}

}

func (env *Env) deviceIndex(c *gin.Context) {
	devices, err := env.SpotifyService.GetDevices()

	if err != nil {
		log.Fatal(err)
		c.String(500, "There was an error")
	}

	c.JSON(200, devices)
}

func (env *Env) deviceSetDefault(c *gin.Context) {
	env.Settings.Set("device.default", c.Param("id"))
	devices, _ := env.SpotifyService.GetDevices()
	var defaultDevice spotify.PlayerDevice

	for _, d := range devices {
		if d.ID.String() == c.Param("id") {
			env.Settings.Set("device.default.json", c.Param("id"))
			defaultDevice = d
			break
		}
	}

	c.JSON(200, defaultDevice)
}

func (env *Env) deviceGetDefault(c *gin.Context) {
	id, _ := env.Settings.Get("device.default")
	c.JSON(200, map[string]interface{}{
		"id": id,
	})
}
