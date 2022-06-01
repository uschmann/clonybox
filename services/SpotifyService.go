package services

import (
	"context"
	"fmt"
	"log"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

type SpotifyService struct {
	Auth          *spotifyauth.Authenticator
	ClientChannel chan *spotify.Client
	State         string
	Client        *spotify.Client
	Settings      *Settings
}

func NewSpotifyService(calbackUrl string, settings *Settings) *SpotifyService {
	return &SpotifyService{
		Auth: spotifyauth.New(
			spotifyauth.WithRedirectURL(calbackUrl),
			spotifyauth.WithScopes(spotifyauth.ScopeUserReadCurrentlyPlaying, spotifyauth.ScopeUserReadPlaybackState, spotifyauth.ScopeUserModifyPlaybackState),
		),
		ClientChannel: make(chan *spotify.Client),
		State:         "abc123",
	}
}

func (s *SpotifyService) StartAuth() string {
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

func (s *SpotifyService) GetDevices() ([]spotify.PlayerDevice, error) {
	return s.Client.PlayerDevices(context.Background())
}

func (s *SpotifyService) GetAuthUrl() string {
	return s.Auth.AuthURL(s.State)
}

func (s *SpotifyService) GetUser() (*spotify.PrivateUser, error) {
	return s.Client.CurrentUser(context.Background())
}
