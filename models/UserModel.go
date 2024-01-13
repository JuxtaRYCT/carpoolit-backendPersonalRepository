package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// User is a representation of a user in the database

type User struct {
	gorm.Model
	Name              string        `gorm:"type:varchar(100);not null;" json:"name" valid:"required~Name is required,matches(^[a-zA-Z ]+$)~Name must be alphabetic"`
	Email             string        `gorm:"type:varchar(100);not null;uniqueIndex" json:"email" valid:"required~Email is required,email~Email is not valid"`
	ProfilePictureURL string        `gorm:"type:text" json:"profile_picture_url" valid:"url~URL is not valid"`
	ContactNumber     string        `gorm:"type:varchar(20);not null" json:"contact_number" valid:"required~Contact number is required,numeric~Contact number must be numeric"`
	Gender            string        `gorm:"type:varchar(10)" json:"gender" valid:"in(male|female|other)~Gender must be male female or other"`
	YOB               uint          `json:"yob" valid:"range(1900|2100)~Year of birth must be between 1900 and 2100"`
	RidesID           pq.Int64Array `gorm:"type:integer[]" json:"ride_ids"`
}

type UserResponse struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	ProfilePictureURL string `json:"profile_picture_url"`
	ContactNumber     string `json:"contact_number"`
	Gender            string `json:"gender"`
	YOB               uint   `json:"yob"`
}
