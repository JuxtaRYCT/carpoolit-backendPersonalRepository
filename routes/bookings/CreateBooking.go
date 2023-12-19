package bookings

import (
	"carpool-backend/database"
	"carpool-backend/helpers"
	"carpool-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateBooking(c *fiber.Ctx) error {
	// Define a variable to hold the Booking struct
	var Booking models.Booking

	// Parse the request body into the Booking struct
	err := c.BodyParser(&Booking)
	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		return c.Status(400).SendString("Error parsing JSON")
	}

	// Validate the Booking struct
	if err := helpers.ValidateBooking(Booking); err != nil {
		log.Printf("Error validating Booking: %v\n", err)
		return c.Status(400).SendString("Error validating Booking")
	}

	// Check if the user has already booked the ride
	existingBooking := models.Booking{}
	if database.Database.Db.Where("user_id = ? AND ride_id = ?", Booking.PassengerID, Booking.RideID).First(&existingBooking).Error == nil {
		log.Println("Similar booking already exists")
		return c.Status(400).SendString("Similar booking already exists")
	}

	// Insert the booking into the database
	result := database.Database.Db.Create(&Booking)
	if result.Error != nil {
		log.Printf("Error creating booking: %v\n", result.Error)
		return c.Status(500).SendString("Error creating booking")
	}

	// Log the booking id
	log.Printf("Booking with id %v created\n", Booking.ID)
	return c.Status(200).JSON(Booking)
}
