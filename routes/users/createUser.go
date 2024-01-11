package users

import (
	"carpool-backend/database"
	"carpool-backend/helpers"
	"carpool-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var createUser models.User

	// Parse and decode the JSON request body into the user struct
	if err := c.BodyParser(&createUser); err != nil {
		log.Printf("Error parsing JSON request body: %v\n", err)
		return c.Status(400).JSON("Error parsing JSON: " + err.Error())
	}

	// validates user, display error if not validated
	if err := helpers.ValidateUser(createUser); err != nil {
		log.Printf("Error validating user: %v\n", err)
		return c.Status(400).SendString("Error validating user")
	}

	// checking if the email already exists in the database
	if err := database.Database.Db.Where("email = ?", createUser.Email).First(&createUser).Error; err == nil {
		log.Printf("Error creating user: Email %v already exists\n", createUser.Email)
		return c.Status(400).SendString("Email already exists")
	}

	// creates a new user in the database
	result := database.Database.Db.Create(&createUser)

	// Check for errors during the user creation in the database
	if result.Error != nil {
		log.Printf("Error creating user: %v\n", result.Error)
		return c.Status(500).SendString("Error creating user")
	}

	// Log success and return the updated user as JSON.
	log.Printf("User successfully created\n")
	return c.Status(200).JSON(createUser)
}
