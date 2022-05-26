package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/uschmann/clonybox/services"
)

func RegisterDeviceRoutes(r *gin.Engine) {

	device := r.Group("/api/device")
	{
		device.GET("/", func(c *gin.Context) {
			devices, err := services.SpotifyService.GetDevices()

			if err != nil {
				log.Fatal(err)
				c.String(500, "There was an error")
			}

			c.JSON(200, devices)
		})

		device.POST("/default/:id", func(c *gin.Context) {
			c.JSON(200, map[string]interface{}{
				"route":  c.FullPath(),
				"method": c.Request.Method,
				"id":     c.Param("id"),
			})
		})
	}

}
