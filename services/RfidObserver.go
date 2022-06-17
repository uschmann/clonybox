package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/uschmann/clonybox/models"
	"github.com/uschmann/clonybox/repos"
	rfidreader "github.com/uschmann/clonybox/services/rfidReader"
	"github.com/zmb3/spotify/v2"
)

type RfidObserver struct {
	RfidChannel        chan rfidreader.RfidEvent
	BroadcastService   *BroadcastService
	PlaybackConfigRepo *repos.PlaybackConfigRepo
	Settings           *Settings
	SpotifyService     *SpotifyService
}

func NewRfidObserver(
	rfidChannel chan rfidreader.RfidEvent,
	broadcastService *BroadcastService,
	playbackConfigRepo *repos.PlaybackConfigRepo,
	setings *Settings,
	spotifyService *SpotifyService,
) *RfidObserver {

	return &RfidObserver{
		RfidChannel:        rfidChannel,
		BroadcastService:   broadcastService,
		PlaybackConfigRepo: playbackConfigRepo,
		Settings:           setings,
		SpotifyService:     spotifyService,
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

				if playbackConfig.SpotifyId != "" && r.IsPlayingPlaybackConfig(playbackConfig) == false {
					r.BroadcastService.Broadcast("playback_config.started", &BroadcastEvent{
						"playback_config": &playbackConfig,
					})
					r.playPlaybackConfig(playbackConfig)
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

func (r *RfidObserver) playPlaybackConfig(playbackConfig *models.PlaybackConfig) {
	if r.Settings.Has("device.default") {
		id, _ := r.Settings.Get("device.default")
		fmt.Println("Play on: " + id)
		deviceId := spotify.ID(id)
		uri := spotify.URI(playbackConfig.SpotifyUrl)

		if playbackConfig.Type == "track" {
			error := r.SpotifyService.Client.PlayOpt(context.Background(), &spotify.PlayOptions{
				DeviceID:       &deviceId,
				URIs:           []spotify.URI{uri},
				PlaybackOffset: &spotify.PlaybackOffset{Position: 0},
			})

			if error != nil {
				fmt.Println(error)
			}
		} else {
			error := r.SpotifyService.Client.PlayOpt(context.Background(), &spotify.PlayOptions{
				DeviceID:        &deviceId,
				PlaybackContext: &uri,
				PlaybackOffset:  &spotify.PlaybackOffset{Position: 0},
			})

			if error != nil {
				fmt.Println(error)
			}
		}
	}
}

func (r *RfidObserver) IsPlayingPlaybackConfig(playbackConfig *models.PlaybackConfig) bool {
	state, err := r.SpotifyService.Client.PlayerState(context.Background())

	if err != nil {
		fmt.Println(err)
		return false
	}

	defaultDeviceId, _ := r.Settings.Get("device.default")
	deviceId := state.Device.ID.String()

	if state.Playing == false {
		return false
	}

	if defaultDeviceId != deviceId {
		return false
	}

	if playbackConfig.Type != r.getContextType(state) {
		return false
	}

	if r.getContextType(state) == "track" {
		return state.Item.ID.String() == playbackConfig.SpotifyId
	} else {
		fmt.Println(r.extractId(string(state.PlaybackContext.URI)), playbackConfig.SpotifyId)
		return r.extractId(string(state.PlaybackContext.URI)) == playbackConfig.SpotifyId
	}
}

func (r *RfidObserver) getContextType(state *spotify.PlayerState) string {
	if state.PlaybackContext.Type == "" {
		return "track"
	}

	return state.PlaybackContext.Type
}

func (r *RfidObserver) extractId(uri string) string {
	return strings.Split(uri, ":")[2]
}
