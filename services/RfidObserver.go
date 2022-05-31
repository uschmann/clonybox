package services

import (
	"fmt"

	"github.com/uschmann/clonybox/repos"
)

type RfidObserver struct {
	RfidChannel        chan string
	BroadcastService   *BroadcastService
	PlaybackConfigRepo *repos.PlaybackConfigRepo
	Settings           *Settings
}

func NewRfidObserver(
	rfidChannel chan string,
	broadcastService *BroadcastService,
	playbackConfigRepo *repos.PlaybackConfigRepo,
	setings *Settings,
) *RfidObserver {

	return &RfidObserver{
		RfidChannel:        rfidChannel,
		BroadcastService:   broadcastService,
		PlaybackConfigRepo: playbackConfigRepo,
		Settings:           setings,
	}
}

func (r *RfidObserver) Observe() {
	for rfid := range r.RfidChannel {
		if r.PlaybackConfigRepo.PlaybackConfigWithRfidExists(rfid) {
			playbackConfig := r.PlaybackConfigRepo.GetPlaybackConfigByRfid(rfid)

			if playbackConfig.SpotifyId != "" {
				r.BroadcastService.Broadcast("playback_config.started", &BroadcastEvent{
					"playback_config": &playbackConfig,
				})

				if r.Settings.Has("device.default") {
					id, _ := r.Settings.Get("device.default")
					fmt.Println("Play on: " + id)
				}
			} else {
				r.BroadcastService.Broadcast("playback_config.scanned", &BroadcastEvent{
					"playback_config": &playbackConfig,
				})
			}
		} else {
			playbackConfig, _ := r.PlaybackConfigRepo.FromRfid(rfid)

			r.BroadcastService.Broadcast("playback_config.scanned", &BroadcastEvent{
				"playback_config": &playbackConfig,
			})
		}
	}
}
