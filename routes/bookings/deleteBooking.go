package bookings

import (
	"carpool-backend/database"
	"carpool-backend/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// DELETE /bookings/:id
// DeleteBooking handles the deletion of a booking
// Accepts a booking id as a url parameter
// Returns a status code and a message
func DeleteBooking(c *fiber.Ctx) error {
	bookingId, err := c.ParamsInt("id")

	if err != nil {
		log.Printf("Error converting booking id to int: %v\n", err)
		return c.Status(400).SendString("Invalid booking id")
	}

	result := database.Database.Db.Delete(&models.Booking{}, bookingId)

	if result.Error != nil {
		log.Printf("Error deleting booking: %v\n", result.Error)
		return c.Status(500).SendString("Could not delete booking")
	}

	if result.RowsAffected == 0 {
		log.Printf("Error deleting Booking, BookingId: %v not found\n", bookingId)
		return c.Status(404).SendString("Booking not found")
	}

	responseMessage := fmt.Sprintf("Booking with id: %v deleted", bookingId)
	log.Println(responseMessage)

	return c.Status(200).SendString(responseMessage)
}
