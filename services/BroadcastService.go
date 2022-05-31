package services

import (
	"encoding/json"
	"fmt"

	"gopkg.in/olahol/melody.v1"
)

type BroadcastService struct {
	melody *melody.Melody
}

type BroadcastEvent map[string]interface{}

func NewBroadcastService(melody *melody.Melody) *BroadcastService {
	return &BroadcastService{melody: melody}
}

type message struct {
	MessageType string          `json:"type"`
	Data        *BroadcastEvent `json:"data"`
}

func (b *BroadcastService) Broadcast(messageType string, data *BroadcastEvent) {
	m := &message{
		MessageType: messageType,
		Data:        data,
	}
	json, err := json.Marshal(m)

	if err != nil {
		fmt.Println(err)
	}

	b.melody.Broadcast(json)
}
