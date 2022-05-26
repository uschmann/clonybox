package handler

import "github.com/gin-gonic/gin"

func RegisterDeviceRoutes(r *gin.Engine) {

	device := r.Group("/api/device")
	{
		device.GET("/", func(c *gin.Context) {
			c.JSON(200, map[string]interface{}{
				"route":  c.FullPath(),
				"method": c.Request.Method,
			})
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
