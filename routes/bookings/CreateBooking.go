package bookings

import (
	"carpool-backend/database"
	"carpool-backend/helpers"
	"carpool-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateBooking(c *fiber.Ctx) error {
	var booking models.Booking

	// Parse the request body into the Booking struct
	err := c.BodyParser(&booking)
	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		return c.Status(400).SendString("Error parsing JSON")
	}

	// Validate the Booking struct
	if err := helpers.ValidateBooking(booking); err != nil {
		log.Printf("Error validating Booking: %v\n", err)
		return c.Status(400).SendString("Error validating Booking")
	}

	// Check if the user has already booked the ride
	existingBooking := models.Booking{}
	if database.Database.Db.Where("user_id = ? AND ride_id = ?", booking.PassengerID, booking.RideID).First(&existingBooking).Error == nil {
		log.Println("Similar booking already exists")
		return c.Status(400).SendString("Similar booking already exists")
	}
	var ride models.Ride
	// Fetch the ride by ID
	if err := database.Database.Db.First(&ride, booking.RideID).Error; err != nil {
		log.Printf("Error finding ride: %v\n", err)
		return c.Status(404).SendString("Ride not found")
	}

	// Check if there are available seats
	if ride.BookedSeats >= ride.TotalSeats {
		log.Println("No available seats for this ride")
		return c.Status(400).SendString("No available seats for this ride")
	}

	// Increment the booked seats count
	if err := database.Database.Db.Model(&ride).Update("booked_seats", ride.BookedSeats+1).Error; err != nil {
		log.Printf("Error updating booked seats for ride: %v\n", err)
		return c.Status(500).SendString("Error updating booked seats for ride")
	}

	result := database.Database.Db.Create(&booking)
	if result.Error != nil {
		log.Printf("Error creating booking: %v\n", result.Error)
		return c.Status(500).SendString("Error creating booking")
	}

	// Log the booking id
	log.Printf("Booking with id %v created\n", booking.ID)
	return c.Status(200).JSON(helpers.CreateResponseBooking(booking))
}
