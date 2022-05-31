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

// Returns true if a PlaybackConfig with the given rfid exists
func (p *PlaybackConfigRepo) PlaybackConfigWithRfidExists(rfid string) bool {
	playbackConfig := &models.PlaybackConfig{}

	result := p.Db.Where("rfid = ?", rfid).Find(playbackConfig)

	if result.RowsAffected > 0 {
		return true
	}

	return false
}

// Creates a new PlaybackConfig with the given rfid
func (p *PlaybackConfigRepo) FromRfid(rfid string) (*models.PlaybackConfig, error) {
	playbackConfig := &models.PlaybackConfig{
		Rfid: rfid,
	}

	result := p.Db.Create(playbackConfig)

	if result.Error != nil {
		return nil, result.Error
	}

	return playbackConfig, nil
}

// Loads a PlaybackConfig by the given Rfid
func (p *PlaybackConfigRepo) GetPlaybackConfigByRfid(rfid string) *models.PlaybackConfig {
	playbackConfig := &models.PlaybackConfig{}

	p.Db.Where("rfid = ?", rfid).First(playbackConfig)

	return playbackConfig
}

func (p *PlaybackConfigRepo) GetPlaybackConfigById(id int) *models.PlaybackConfig {
	playbackConfig := &models.PlaybackConfig{}

	p.Db.Where("id = ?", id).First(playbackConfig)

	return playbackConfig
}

// Updates the given PlaylistConfig
func (p *PlaybackConfigRepo) UpdatePlaybackConfig(playbackConfig *models.PlaybackConfig) {
	p.Db.Save(playbackConfig)
}
