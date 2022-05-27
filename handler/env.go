package handler

import (
	"github.com/uschmann/clonybox/services"
	"gorm.io/gorm"
)

type Env struct {
	Db             *gorm.DB
	SpotifyService *services.SpotifyService
	Settings       *services.Settings
}
