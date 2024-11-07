package database

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string
	Name           string
	IDPUserId      string `gorm:"column:idp_user_id;uniqueIndex"`
	IDPType        string `gorm:"column:idp_type"`
	ProfilePicture string
	IsRyan         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type InvitationCode struct {
	gorm.Model
	Code        string `gorm:"uniqueIndex"`
	ClaimedDate time.Time
	Amount      int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Transaction struct {
	gorm.Model
	SenderID    string
	RecipientID string
	Amount      int
	Description string
	CodeID      uint
	Code        InvitationCode
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var DB *gorm.DB

func InitDatabase() error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
	db.AutoMigrate(&InvitationCode{})
	db.AutoMigrate(&Transaction{})

	DB = db

	return nil
}
