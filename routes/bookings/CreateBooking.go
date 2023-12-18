package bookings

import (
	"carpool-backend/database"
	"carpool-backend/helpers"
	"carpool-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateBooking(c *fiber.Ctx) error {

	var Booking models.Booking
	err := c.BodyParser(&Booking)

	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		return c.Status(400).SendString("Error parsing JSON")
	}

	if err := helpers.ValidateBooking(Booking); err != nil {
		log.Printf("Error validating Booking: %v\n", err)
		return c.Status(400).SendString("Error validating Booking")
	}
	existingBooking := models.Booking{}
	if database.Database.Db.Where("user_id = ? AND ride_id = ?", Booking.PassengerID, Booking.RideID).First(&existingBooking).Error == nil {
		log.Println("Similar booking already exists")
		return c.Status(400).SendString("Similar booking already exists")
	}

	result := database.Database.Db.Create(&Booking)

	if result.Error != nil {
		log.Printf("Error creating booking: %v\n", result.Error)
		return c.Status(500).SendString("Error creating booking")
	}

	log.Printf("Booking with id %v created\n", Booking.ID)
	return c.Status(200).JSON(Booking)
}
