package storage

import (
	"fmt"
	"log"

	"github.com/uschmann/clonybox/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDb(filename string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db.AutoMigrate(&models.Setting{})

	fmt.Println("DB opened successfully")

	return db, nil
}
