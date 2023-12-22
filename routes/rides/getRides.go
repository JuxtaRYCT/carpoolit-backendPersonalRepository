package rides

import (
	"carpool-backend/database"
	"carpool-backend/helpers"
	"carpool-backend/models"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetRides handles fetching all rides from the database which are not full and are yet to start
// Returns an HTTP status code and a message indicating the result of the operation
//
// Parameters:
//   - c: Fiber context
//
// Returns:
//   - error: An error message indicating the result of the operation and an HTTP status code
func GetRides(c *fiber.Ctx) error {
	var rides []models.Ride

	result := database.Database.Db.Where("booked_seats < total_seats AND start_time > NOW()").Find(&rides)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Printf("No rides found\n")
		return c.Status(404).SendString("No rides found")
	}

	if result.Error != nil {
		log.Printf("Error fetching rides: %v\n", result.Error)
		return c.Status(500).SendString("Error fetching rides")
	}

	var responseRides []models.RideResponse

	for _, ride := range rides {
		responseRides = append(responseRides, helpers.CreateResponseRide(ride))
	}

	return c.Status(200).JSON(responseRides)
}

func GetRidesById(c *fiber.Ctx) error {
	var ride models.Ride

	// Fetch the ride by ID
	if err := database.Database.Db.First(&ride, c.Params("id")).Error; err != nil {
		log.Printf("Error finding ride: %v\n", err)
		return c.Status(404).SendString("Ride not found")
	}

	return c.Status(200).JSON(helpers.CreateResponseRide(ride))
}
