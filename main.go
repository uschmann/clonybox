package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uschmann/clonybox/handler"
	"github.com/uschmann/clonybox/services"
	"github.com/uschmann/clonybox/storage"
)

func main() {
	r := gin.Default()
	storage.OpenDb("test.db")

	handler.RegisterDeviceRoutes(r)
	handler.RegisterAuthRoutes(r)

	services.SpotifyService.StartAuth()

	fmt.Println(storage.Get("ef"))

	r.Run()
}
