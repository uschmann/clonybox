package models

import "gorm.io/datatypes"

type PlaybackConfig struct {
	ID         uint           `json:"id"`
	Rfid       string         `json:"rfid"`
	Name       string         `json:"name"`
	SpotifyUrl string         `json:"spotify_url"`
	SpotifyId  string         `json:"spotify_id"`
	MetaData   datatypes.JSON `json:"metadata"`
	Type       string         `json:"type"`
}
