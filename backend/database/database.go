package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string
	Name           string
	IDPUserId      string `gorm:"column:idp_user_id"`
	IDPType        string `gorm:"column:idp_type"`
	ProfilePicture string
}

var DB *gorm.DB

func InitDatabase() error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	DB = db

	return nil
}
