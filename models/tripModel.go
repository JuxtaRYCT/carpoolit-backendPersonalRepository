package models

import (
	"gorm.io/gorm"
	"time"
)

type Trip struct {
	gorm.Model
	ID uint  `gorm:"primaryKey" json:"id"`
	StartLoc string `json:"startLoc"`
	EndLoc string `json:"endLoc"`
	StartTime time.Time `json:"startTime"`
	TotalCost float64 `json:"totalCost"`
	PassengerIDs []uint `gorm:"many2many:trip_passengers;" json:"passengerIDs"`
	HostID uint `json:"hostID"`
}