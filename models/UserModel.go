package models

import "gorm.io/gorm"

// User is a representation of a user in the database
type User struct {
	gorm.Model
	ID                uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name              string `gorm:"size:255;not null;" json:"name"`
	Email             string `gorm:"size:255;not null;uniqueIndex" json:"email"` // Added uniqueIndex for better query performance
	ProfilePictureURL string `gorm:"size:255" json:"profile_picture_url"`        // Removed 'not null' if this field can be optional
}
