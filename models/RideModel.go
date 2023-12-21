package models

import (
	"time"

	"gorm.io/gorm"
)

// Ride is a representation of a ride in the database
type Ride struct {
	gorm.Model
	HostUserID    uint      `gorm:"not null" json:"host_user_id" valid:"required~Host user ID is required"`
	HostUser      User      `gorm:"foreignKey:HostUserID;references:ID" json:"host_user"`
	StartLocation string    `gorm:"size:255;not null;" json:"start_location" valid:"required~Start location is required"`
	EndLocation   string    `gorm:"size:255;not null;" json:"end_location" valid:"required~End location is required"`
	StartTime     time.Time `gorm:"not null;" json:"start_time" valid:"required~Start time is required"`
	TotalSeats    uint      `gorm:"not null;" json:"total_seats" valid:"required~Total seats is required"`
	BookedSeats   uint      `gorm:"not null;" json:"booked_seats" valid:"required~Booked seats is required"`
	TotalPrice    uint      `gorm:"not null;" json:"total_price" valid:"required~Total price is required"`
}

type RideResponse struct {
	HostUserID    uint      `json:"host_user_id"`
	StartLocation string    `json:"start_location"`
	EndLocation   string    `json:"end_location"`
	StartTime     time.Time `json:"start_time"`
	TotalSeats    uint      `json:"total_seats"`
	BookedSeats   uint      `json:"booked_seats"`
	TotalPrice    uint      `json:"total_price"`
}
