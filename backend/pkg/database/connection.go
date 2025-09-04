package database

import (
	"github.com/bhati00/workova/backend/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg config.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(cfg.DBpath), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = db
	return db
}
