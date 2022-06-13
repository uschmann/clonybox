package handler

import (
	"github.com/uschmann/clonybox/config"
	"github.com/uschmann/clonybox/repos"
	"github.com/uschmann/clonybox/services"
	rfidreader "github.com/uschmann/clonybox/services/rfidReader"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
)

type Env struct {
	Config             *config.Config
	Db                 *gorm.DB
	SpotifyService     *services.SpotifyService
	Settings           *services.Settings
	PlaybackConfigRepo *repos.PlaybackConfigRepo
	RfidChannel        chan rfidreader.RfidEvent
	Melody             *melody.Melody
	BroadcastService   *services.BroadcastService
	RfidObserver       *services.RfidObserver
}
