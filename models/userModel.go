package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Phone int `json:"phone"`
	Email string `json:"email"`
}