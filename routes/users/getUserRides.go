package users

import (
	"carpool-backend/database"
	"carpool-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

type RideResponse struct {
	ID            uint   `json:"ID"`
	CreatedAt     string `json:"CreatedAt"`
	UpdatedAt     string `json:"UpdatedAt"`
	HostUserID    uint   `json:"host_user_id"`
	StartLocation string `json:"start_location"`
	EndLocation   string `json:"end_location"`
	StartTime     string `json:"start_time"`
	TotalSeats    uint   `json:"total_seats"`
	BookedSeats   uint   `json:"booked_seats"`
	TotalPrice    uint   `json:"total_price"`
	RideStatus    string `json:"ride_status"`
}

func GetUserRides(c *fiber.Ctx) error {
	UserID := c.Params("userID")

	var bookings []models.Booking

	// Retrieve bookings made by the user
	if err := database.Database.Db.Where("passenger_id = ?", UserID).Find(&bookings).Error; err != nil {
		log.Printf("Error finding bookings: %v\n", err)
		return c.Status(500).SendString("Error finding bookings")
	}

	// Extract unique ride IDs from bookings
	rideIDs := make([]uint, len(bookings))
	for i, booking := range bookings {
		rideIDs[i] = booking.RideID
	}

	// Retrieve rides associated with the extracted ride IDs
	var rides []models.Ride
	if err := database.Database.Db.Where("id IN (?)", rideIDs).Find(&rides).Error; err != nil {
		log.Printf("Error finding rides: %v\n", err)
		return c.Status(500).SendString("Error finding rides")
	}

	// Create the response
	var ridesResponse []RideResponse
	for _, ride := range rides {
		rideResponse := RideResponse{
			ID:            ride.ID,
			CreatedAt:     ride.CreatedAt.String(),
			UpdatedAt:     ride.UpdatedAt.String(),
			HostUserID:    ride.HostUserID,
			StartLocation: ride.StartLocation,
			EndLocation:   ride.EndLocation,
			StartTime:     ride.StartTime.String(),
			TotalSeats:    ride.TotalSeats,
			BookedSeats:   ride.BookedSeats,
			TotalPrice:    ride.TotalPrice,
			RideStatus:    ride.RideStatus,
		}
		ridesResponse = append(ridesResponse, rideResponse)
	}

	return c.Status(200).JSON(ridesResponse)
}
