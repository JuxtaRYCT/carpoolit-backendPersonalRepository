package models

import (
	"gorm.io/gorm"
)

// Booking is a representation of a booking in the database
type Booking struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement" json:"id"`
	RideID      uint `gorm:"not null;uniqueIndex:idx_booking_ride_passenger" json:"ride_id"`
	Ride        Ride `gorm:"foreignKey:RideID;references:ID" json:"ride"`
	PassengerID uint `gorm:"not null;uniqueIndex:idx_booking_ride_passenger" json:"passenger_id"`
	Passenger   User `gorm:"foreignKey:PassengerID;references:ID" json:"passenger"`
}
