package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uschmann/clonybox/handler"
	"github.com/uschmann/clonybox/services"
)

func main() {
	r := gin.Default()

	handler.RegisterDeviceRoutes(r)
	handler.RegisterAuthRoutes(r)

	services.SpotifyService.StartAuth()

	r.Run()
}
