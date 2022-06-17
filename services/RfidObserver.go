package services

import (
	"fmt"

	"github.com/uschmann/clonybox/repos"
	rfidreader "github.com/uschmann/clonybox/services/rfidReader"
)

type RfidObserver struct {
	RfidChannel        chan rfidreader.RfidEvent
	BroadcastService   *BroadcastService
	PlaybackConfigRepo *repos.PlaybackConfigRepo
	Settings           *Settings
	SpotifyService     *SpotifyService
	AudioPlayer        *AudioPlayer
}

func NewRfidObserver(
	rfidChannel chan rfidreader.RfidEvent,
	broadcastService *BroadcastService,
	playbackConfigRepo *repos.PlaybackConfigRepo,
	setings *Settings,
	spotifyService *SpotifyService,
	audioPlayer *AudioPlayer,
) *RfidObserver {

	return &RfidObserver{
		RfidChannel:        rfidChannel,
		BroadcastService:   broadcastService,
		PlaybackConfigRepo: playbackConfigRepo,
		Settings:           setings,
		SpotifyService:     spotifyService,
		AudioPlayer:        audioPlayer,
	}
}

func (r *RfidObserver) Observe() {
	for event := range r.RfidChannel {
		fmt.Println(event)
		rfid := event.Rfid

		switch event.Type {
		case rfidreader.EVENT_SCANNED:
			if r.PlaybackConfigRepo.PlaybackConfigWithRfidExists(rfid) {
				playbackConfig := r.PlaybackConfigRepo.GetPlaybackConfigByRfid(rfid)

				if playbackConfig.SpotifyId != "" && r.AudioPlayer.IsPlayingPlaybackConfig(playbackConfig) == false {
					r.BroadcastService.Broadcast("playback_config.started", &BroadcastEvent{
						"playback_config": &playbackConfig,
					})
					r.AudioPlayer.playPlaybackConfig(playbackConfig)
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
}
