package repos

import (
	"github.com/uschmann/clonybox/models"
	"gorm.io/gorm"
)

type PlaybackConfigRepo struct {
	Db *gorm.DB
}

func NewPlaybackConfigRepo(db *gorm.DB) *PlaybackConfigRepo {
	return &PlaybackConfigRepo{Db: db}
}

// Load all PlaybackConfigs from the database
func (p *PlaybackConfigRepo) GetAllPlaybackConfigs() ([]models.PlaybackConfig, error) {
	var playbackConfigs []models.PlaybackConfig
	result := p.Db.Find(&playbackConfigs)

	if result.Error != nil {
		return nil, result.Error
	}

	return playbackConfigs, nil
}

// Stores a new PlaybackConfig in the database
func (p *PlaybackConfigRepo) CreatePlaybackConfig(playbackConfig *models.PlaybackConfig) (*models.PlaybackConfig, error) {
	result := p.Db.Create(playbackConfig)

	if result.Error != nil {
		return nil, result.Error
	}

	return playbackConfig, nil
}
