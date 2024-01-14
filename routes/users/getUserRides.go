package users

import (
	"carpool-backend/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

type RideResponse struct {
	ID            uint   `json:"ID"`
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

	var ridesResponse []RideResponse

	// Use GORM's Select to manually specify the fields for the response structure
	if err := database.Database.Db.
		Table("rides").
		Select("rides.id, rides.host_user_id, rides.start_location, rides.end_location, rides.start_time, rides.total_seats, rides.booked_seats, rides.total_price, rides.ride_status").
		Joins("JOIN bookings ON rides.id = bookings.ride_id").
		Where("bookings.passenger_id = ?", UserID).
		Scan(&ridesResponse).Error; err != nil {
		log.Printf("Error finding rides: %v\n", err)
		return c.Status(500).SendString("Error finding rides")
	}

	return c.Status(200).JSON(ridesResponse)
}
