package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func OpenDb(filename string) {
	Db, err = gorm.Open(sqlite.Open(filename), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return
	}

	Db.AutoMigrate(&Setting{})

	fmt.Println("DB opened successfully")
}
