package models

import (
	"gorm.io/gorm"
)

// Booking is a representation of a booking in the database
type Booking struct {
	gorm.Model
	RideID      uint `gorm:"not null;uniqueIndex:idx_booking_ride_passenger" json:"ride_id" valid:"required~Ride ID is required"`
	Ride        Ride `gorm:"foreignKey:RideID;references:ID" json:"ride"`
	PassengerID uint `gorm:"not null;uniqueIndex:idx_booking_ride_passenger" json:"passenger_id" valid:"required~Passenger ID is required"`
	Passenger   User `gorm:"foreignKey:PassengerID;references:ID" json:"passenger"`
}

type BookingResponse struct {
	RideID      uint `json:"ride_id"`
	PassengerID uint `json:"passenger_id"`
}
