package rides

import (
	"carpool-backend/database"
	"carpool-backend/helpers"
	"carpool-backend/models"
	"log"
	"strconv"

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
	id, err := strconv.Atoi(c.Params("id")) // used strconv.Atoi
	var ride models.Ride
	if err != nil || id <= 0 { // Check provided id is correct or not, displays error if incorrect id
		log.Printf("Invalid or missing ride ID: %v", err)
		return c.Status(400).JSON("Please provide valid ID")
	}
	if err := database.Database.Db.Find(&ride, id).Error; err != nil { // Find ride via id, displays error if ride not fetched
		log.Printf("Error fetching the ride with ID %d: %v", id, err)
		return c.Status(400).JSON("Error fetching the ride: " + err.Error())
	}
	// Parse and decode the JSON request body into the ride struct
	if err := c.BodyParser(&ride); err != nil {
		log.Printf("Error parsing JSON request body: %v", err)
		return c.Status(400).JSON("Error parsing JSON: " + err.Error())
	}
	if err := helpers.ValidateRide(ride); err != nil { // validates ride, display error if not validated
		log.Printf("Error validating ride: %v\n", err)
		return c.Status(400).SendString("Error validating ride")
	}
	// Save the ride to the database.
	updatedRide := database.Database.Db.Save(&ride)
	// Check for errors during the database update
	if updatedRide.Error != nil {
		log.Printf("Error updating ride with ID %d: %v", ride.ID, updatedRide.Error)
		return c.Status(500).SendString("Error updating ride")
	}
	// Check if any rows were affected during the update
	if updatedRide.RowsAffected == 0 {
		log.Printf("Ride with ID %d not found", ride.ID)
		return c.Status(500).SendString("Ride not found")
	}
	// Log success and return the updated ride as JSON.
	log.Printf("Ride with ID %d has been updated", ride.ID)
	return c.Status(200).JSON(ride)
}
