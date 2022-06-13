package handler

import (
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterSettingsHandler(r *gin.Engine, env *Env) {
	group := r.Group("/api/settings")
	group.POST("/reset", env.resetDevice)
}

// Deletes all data and restarts
func (env *Env) resetDevice(c *gin.Context) {
	os.Remove(env.Config.DatabaseFile)

	c.JSON(200, &gin.H{
		"status": "OK",
	})

	os.Exit(0)
}
