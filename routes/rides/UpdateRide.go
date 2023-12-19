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
		return c.Status(400).JSON("Please provide valid ID")
	}
	if err := database.Database.Db.Find(&ride, id).Error; err != nil { // Find ride via id, displays error if ride not fetched
		return c.Status(400).JSON("Error fetching the ride: " + err.Error())
	}
	if err := helpers.ValidateRide(ride); err != nil { // validates ride, display error if not validated
		log.Printf("Error validating ride: %v\n", err)
		return c.Status(400).SendString("Error validating ride")
	}
	type UpdateRide struct { // struct for ride update
		HostUser      models.User `json:"host_user"`
		StartLocation string      `json:"start_location"`
		EndLocation   string      `json:"end_location"`
		TotalSeats    uint        `json:"total_seats"`
		BookedSeats   uint        `json:"booked_seats"`
		TotalPrice    uint        `json:"total_price"`
	}
	var updateData UpdateRide // creates variable updateData to store data of UpdateRide

	if err := c.BodyParser(&updateData); err != nil { // parse the request body and decode it into the updateData variable, display error if unable to do so.
		return c.Status(500).JSON(err.Error())
	}
	// Update ride properties with new data.
	ride.HostUser = updateData.HostUser
	ride.StartLocation = updateData.StartLocation
	ride.EndLocation = updateData.EndLocation
	ride.TotalSeats = updateData.TotalSeats
	ride.BookedSeats = updateData.BookedSeats
	ride.TotalPrice = updateData.TotalPrice

	// Save the updated ride to the database.
	updatedRide := database.Database.Db.Save(&ride)
	if updatedRide.RowsAffected == 0 {
		return c.Status(500).SendString("No ride updated")
	}

	// Log success and return the updated ride as JSON.
	log.Printf("Ride with ID %d has been updated", ride.ID)
	return c.Status(200).JSON(ride)
}
