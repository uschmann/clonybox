package handler

import (
	"github.com/uschmann/clonybox/repos"
	"github.com/uschmann/clonybox/services"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
)

type Env struct {
	Db                 *gorm.DB
	SpotifyService     *services.SpotifyService
	Settings           *services.Settings
	PlaybackConfigRepo *repos.PlaybackConfigRepo
	RfidChannel        chan string
	Melody             *melody.Melody
	BroadcastService   *services.BroadcastService
	RfidObserver       *services.RfidObserver
}
