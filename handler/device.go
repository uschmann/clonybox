package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterDeviceRoutes(r *gin.Engine, env *Env) {

	device := r.Group("/api/device")
	{
		device.GET("/", env.deviceIndex)
		device.POST("/default/:id", env.deviceDefault)
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

func (env *Env) deviceDefault(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"route":  c.FullPath(),
		"method": c.Request.Method,
		"id":     c.Param("id"),
	})
}
