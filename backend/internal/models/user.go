package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Email     string         `gorm:"uniqueIndex;not null"`
	Username  string         `gorm:"uniqueIndex;not null"`
	Password  string         `gorm:"not null"` // Hashed password
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// MigrateUsers applies the database migration for the User model
func MigrateUsers(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
