package rides

import (
	"carpool-backend/database"
	"carpool-backend/helpers"
	"carpool-backend/models"
	"log"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateRide handles the creation of a ride in the database
// Accepts a JSON object containing the ride information
// Returns an HTTP status code and a message indicating the result of the operation
//
// Parameters:
//   - c: Fiber context
//
// Returns:
//   - error: An error message indicating the result of the operation and an HTTP status code

func ValidateRide(ride models.Ride) error {
	_, err := govalidator.ValidateStruct(ride)
	return err
}

func CreateRide(c *fiber.Ctx) error {
	var ride models.Ride

	//REQUEST VALIDATION
	//Logs a 405 error if the request method is not POST
	if c.Method() != "POST" {
		log.Printf("Request method is not POST\n")
		return c.Status(405).SendString("Request method is not POST")
	}

	//Logs a 400 error if the JSON request is invalid
	err := c.BodyParser(&ride) // Parse the request body into a Ride struct
	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		return c.Status(400).SendString("Error parsing JSON")
	}

	//Logs a 400 error if the Ride struct is invalid
	err = ValidateRide(ride) // Validate the Ride struct
	if err != nil {
		log.Printf("Error validating ride: %v\n", err)
		return c.Status(400).SendString("Error validating ride")
	}

	//DATA VALIDATION
	//Checks if the host user ID exists
	var hostUser models.User
	result := database.Database.Db.First(&hostUser, ride.HostUserID)
	if result.Error == gorm.ErrRecordNotFound {
		log.Printf("Host user ID does not exist")
		return c.Status(400).SendString("Host user ID does not exist")
	} else if result.Error != nil {
		log.Printf("Error finding host user: %v\n", result.Error)
		return c.Status(502).SendString("Error finding host user")
	}

	//Checks if start time is in the future
	if ride.StartTime.Before(time.Now()) {
		log.Printf("Start time is in the past")
		return c.Status(400).SendString("Start time is in the past")
	}

	//Checks if the host user has enough seats
	if ride.TotalSeats <= ride.BookedSeats {
		log.Printf("Total seats available should be more than booked seats")
		return c.Status(400).SendString("Total seats available should be more than booked seats")
	}

	//Obtaining the user from the db using the host user ID
	ride.HostUser = hostUser

	//CREATING THE RIDE
	result = database.Database.Db.Create(&ride)

	if result.Error != nil {
		log.Printf("Error creating ride: %v\n", result.Error)
		return c.Status(500).SendString("Error creating ride")
	}

	log.Printf("Ride with id %v created\n", ride.ID)
	return c.Status(200).JSON(helpers.CreateResponseRide(ride))
}
