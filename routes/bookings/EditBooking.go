package bookings

import (
	"carpool-backend/database"
	"carpool-backend/helpers"
	"carpool-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func EditBooking(c *fiber.Ctx) error {
	var booking models.Booking

	// Parse and decode the JSON request body into the booking struct
	if err := c.BodyParser(&booking); err != nil {
		log.Printf("Error parsing JSON request body: %v\n", err)
		return c.Status(400).JSON("Error parsing JSON: " + err.Error())
	}

	// Validate the booking, display an error if not validated
	if err := helpers.ValidateBooking(booking); err != nil {
		log.Printf("Error validating booking: %v\n", err)
		return c.Status(400).SendString("Error validating booking")
	}

	// Update the booking in the database.
	result := database.Database.Db.Model(&models.Booking{}).Where("id = ?", booking.ID).Updates(&booking)

	// Check for errors during the database update
	if result.Error != nil {
		log.Printf("Error updating booking: %v\n", result.Error)
		return c.Status(500).SendString("Error updating booking")
	}

	// Check if any rows were affected during the update
	if result.RowsAffected == 0 {
		log.Printf("Booking with ID %v not found\n", booking.ID)
		return c.Status(404).SendString("Booking not found")
	}

	// Log success and return the updated booking as JSON.
	log.Printf("Booking with ID %v has been updated", booking.ID)
	return c.Status(200).JSON(helpers.CreateResponseBooking(booking))
}
