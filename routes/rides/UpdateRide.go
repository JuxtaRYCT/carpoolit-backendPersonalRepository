package rides

import (
	"carpool-backend/database"
	"carpool-backend/helpers"
	"carpool-backend/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// UpdateRide handles updating a ride based on the provided ID in the Fiber context.
// Accepts a JSON object containing the ride information
// Returns an HTTP status code and a message indicating the result of the operation
//
// Parameters:
//   - c: Fiber context
//
// Returns:
//   - error: An error message indicating the result of the operation and an HTTP status code

func UpdateRide(c *fiber.Ctx) error {
	var updatedRide models.Ride

	// Parse and decode the JSON request body into the ride struct
	if err := c.BodyParser(&updatedRide); err != nil {
		log.Printf("Error parsing JSON request body: %v\n", err)
		return c.Status(400).JSON("Error parsing JSON: " + err.Error())
	}

	// validates ride, display error if not validated
	if err := helpers.ValidateRide(updatedRide); err != nil {
		log.Printf("Error validating ride: %v\n", err)
		return c.Status(400).SendString("Error validating ride")
	}

	// Update the ride to the database, using id.
	result := database.Database.Db.Model(&models.Ride{}).Where("id = ?", updatedRide.ID).Updates(&updatedRide)

	// Check for errors during the database update
	if result.Error != nil {
		log.Printf("Error updating ride: %v\n", result.Error)
		return c.Status(500).SendString("Error updating ride")
	}

	// Check if any rows were affected during the update
	if result.RowsAffected == 0 {
		log.Printf("Ride with ID %v not found\n", updatedRide.ID)
		return c.Status(404).SendString("Ride not found")
	}

	// Log success and return the updated ride as JSON.
	log.Printf("Ride with ID %v successfully updated\n", updatedRide.ID)
	return c.Status(200).JSON(helpers.CreateResponseRide(updatedRide))
}

func UpdateRideStatus(c *fiber.Ctx) error {
	var updatedRide struct {
		ID        uint      `json:"id" validate:"required"`
		StartTime time.Time `gorm:"not null;" json:"start_time" valid:"required~Start time is required"`
	}

	// Parse and decode the JSON request body into the ride struct
	if err := c.BodyParser(&updatedRide); err != nil {
		log.Printf("Error parsing JSON request body: %v\n", err)
		return c.Status(400).JSON("Error parsing JSON: " + err.Error())
	}

	// Check if the current time has passed the ride start time
	currentTime := time.Now().UTC()
	if currentTime.Before(updatedRide.StartTime) {
		log.Printf("Cannot update ride status before the ride start time")
		return c.Status(400).SendString("Cannot update ride status before the ride start time")
	}

	// Update the ride status in the database, using id.
	result := database.Database.Db.Model(&models.Ride{}).Where("id = ?", updatedRide.ID).Update("ride_status", "completed")

	// Check for errors during the database update
	if result.Error != nil {
		log.Printf("Error updating ride: %v\n", result.Error)
		return c.Status(500).SendString("Error updating ride")
	}

	// Check if any rows were affected during the update
	if result.RowsAffected == 0 {
		log.Printf("Ride with ID %v not found\n", updatedRide.ID)
		return c.Status(404).SendString("Ride not found")
	}

	// Log success and return the updated ride as JSON.
	log.Printf("Ride with ID %v successfully updated\n", updatedRide.ID)
	return c.Status(200).JSON("Ride Status Updated")
}
