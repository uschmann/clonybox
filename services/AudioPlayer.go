package services

import (
	"context"
	"fmt"

	"github.com/uschmann/clonybox/models"
	"github.com/zmb3/spotify/v2"
)

type AudioPlayer struct {
	Settings       *Settings
	SpotifyService *SpotifyService
}

func NewAudioPlayer(settings *Settings, spotifyService *SpotifyService) *AudioPlayer {
	return &AudioPlayer{
		Settings:       settings,
		SpotifyService: spotifyService,
	}
}

func (a *AudioPlayer) playPlaybackConfig(playbackConfig *models.PlaybackConfig) {
	if a.Settings.Has("device.default") {
		id, _ := a.Settings.Get("device.default")
		fmt.Println("Play on: " + id)
		deviceId := spotify.ID(id)
		uri := spotify.URI(playbackConfig.SpotifyUrl)

		if playbackConfig.Type == "track" {
			error := a.SpotifyService.Client.PlayOpt(context.Background(), &spotify.PlayOptions{
				DeviceID:       &deviceId,
				URIs:           []spotify.URI{uri},
				PlaybackOffset: &spotify.PlaybackOffset{Position: 0},
			})

			if error != nil {
				fmt.Println(error)
			}
		} else {
			error := a.SpotifyService.Client.PlayOpt(context.Background(), &spotify.PlayOptions{
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

func (a *AudioPlayer) IsPlayingPlaybackConfig(playbackConfig *models.PlaybackConfig) bool {
	state, err := a.SpotifyService.Client.PlayerState(context.Background())

	if err != nil {
		fmt.Println(err)
		return false
	}

	defaultDeviceId, _ := a.Settings.Get("device.default")
	deviceId := state.Device.ID.String()

	if state.Playing == false {
		return false
	}

	if defaultDeviceId != deviceId {
		return false
	}

	contextType := a.SpotifyService.GetContextType(state)
	if playbackConfig.Type != contextType {
		return false
	}

	if contextType == "track" {
		return state.Item.ID.String() == playbackConfig.SpotifyId
	} else {
		return a.SpotifyService.ExtractId(string(state.PlaybackContext.URI)) == playbackConfig.SpotifyId
	}
}
