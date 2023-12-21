package models

import "gorm.io/gorm"

// User is a representation of a user in the database
type User struct {
	gorm.Model
	Name              string `gorm:"size:255;not null;" json:"name" valid:"required~Name is required"`
	Email             string `gorm:"size:255;not null;uniqueIndex" json:"email" valid:"required~Email is required"` // Added uniqueIndex for better query performance
	ProfilePictureURL string `gorm:"size:255" json:"profile_picture_url"`                                           // Removed 'not null' if this field can be optional
}
