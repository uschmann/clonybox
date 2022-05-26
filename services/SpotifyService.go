package services

import (
	"context"
	"fmt"
	"log"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

type spotifyService struct {
	Auth          *spotifyauth.Authenticator
	ClientChannel chan *spotify.Client
	State         string
	Client        *spotify.Client
}

var SpotifyService *spotifyService = &spotifyService{
	Auth: spotifyauth.New(
		spotifyauth.WithRedirectURL("http://localhost:8080/callback"),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadCurrentlyPlaying, spotifyauth.ScopeUserReadPlaybackState, spotifyauth.ScopeUserModifyPlaybackState),
	),
	ClientChannel: make(chan *spotify.Client),
	State:         "abc123",
}

func (s *spotifyService) StartAuth() string {
	url := s.Auth.AuthURL(s.State)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	go func() {
		// wait for auth to complete
		s.Client = <-s.ClientChannel

		// use the client to make calls that require authorization
		user, err := s.Client.CurrentUser(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("You are logged in as:", user.ID)
	}()

	return url
}

func (s *spotifyService) GetDevices() ([]spotify.PlayerDevice, error) {
	return s.Client.PlayerDevices(context.Background())
}
