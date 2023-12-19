package models

import (
	"gorm.io/gorm"
	"time"
)

// Ride is a representation of a ride in the database
type Ride struct {
	gorm.Model
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	HostUserID    uint      `gorm:"not null" json:"host_user_id"`
	HostUser      User      `gorm:"foreignKey:HostUserID;references:ID" json:"host_user"`
	StartLocation string    `gorm:"size:255;not null;" json:"start_location"`
	EndLocation   string    `gorm:"size:255;not null;" json:"end_location"`
	StartTime     time.Time `gorm:"not null;" json:"start_time"`
	TotalSeats    uint      `gorm:"not null;" json:"total_seats"`
	BookedSeats   uint      `gorm:"not null;" json:"booked_seats"`
	TotalPrice    uint      `gorm:"not null;" json:"total_price"`
}
