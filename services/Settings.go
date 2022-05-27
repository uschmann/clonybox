package services

import (
	"github.com/uschmann/clonybox/models"
	"gorm.io/gorm"
)

type Settings struct {
	db *gorm.DB
}

func NewSettings(db *gorm.DB) *Settings {
	return &Settings{db: db}
}

func (s *Settings) Set(key, value string) {
	if s.Has(key) {
		s.db.Model(&models.Setting{}).Where("Key = ?", key).Update("Value", value)
	} else {
		s.db.Create(&models.Setting{
			Key:   key,
			Value: value,
		})
	}
}

func (s *Settings) Has(key string) bool {
	setting := models.Setting{}

	result := s.db.Where("Key = ?", key).Find(&setting)

	if result.RowsAffected > 0 {
		return true
	}

	return false
}

func (s *Settings) Get(key string) (string, bool) {
	if !s.Has(key) {
		return "", false
	}

	setting := models.Setting{}

	s.db.Where("Key = ?", key).First(&setting)

	return setting.Value, true
}
