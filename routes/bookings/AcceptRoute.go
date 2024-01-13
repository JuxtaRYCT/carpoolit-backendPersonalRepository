package bookings

import (
	"carpool-backend/database"
	"carpool-backend/helpers"
	"carpool-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func AcceptRoute(c *fiber.Ctx) error {
	bookingID := c.Params("bookingID")

	var booking models.Booking

	if err := database.Database.Db.First(&booking, bookingID).Error; err != nil {
		log.Printf("Booking does not exist")
		return c.Status(404).SendString("Booking not found")
	}

	if err := database.Database.Db.Model(&booking).Update("request_status", "accepted").Error; err != nil {
		log.Printf("Error updating booking status: %v\n", err)
		return c.Status(500).SendString("Error updating booking status")
	}

	var ride models.Ride
	RideID := booking.RideID
	if err := database.Database.Db.First(&ride, RideID).Error; err != nil {
		log.Printf("Ride not Found")
		return c.Status(404).SendString("Ride not Found")
	}

	if ride.BookedSeats >= ride.TotalSeats {
		log.Println("No available seats for this ride")
		return c.Status(400).SendString("No available seats for this ride")
	}

	// Increment the booked seats count
	if err := database.Database.Db.Model(&ride).Update("booked_seats", ride.BookedSeats+1).Error; err != nil {
		log.Printf("Error updating booked seats for ride: %v\n", err)
		return c.Status(500).SendString("Error updating booked seats for ride")
	}

	var User models.User
	UserID := booking.PassengerID

	if err := database.Database.Db.First(&User, UserID).Error; err != nil {
		log.Printf("User not Found")
		return c.Status(404).SendString("User not Found")
	}

	// Append ride ID to user's past rides array
	User.RidesID = append(User.RidesID, int64(ride.ID))

	if err := database.Database.Db.Save(&User).Error; err != nil {
		log.Printf("Error updating user's past rides: %v\n", err)
		return c.Status(500).SendString("Error updating user's past rides")
	}

	return c.Status(200).JSON(helpers.CreateResponseBooking(booking))

}
