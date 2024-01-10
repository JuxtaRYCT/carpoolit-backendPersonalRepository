package users

import (
	"carpool-backend/database"
	"carpool-backend/helpers"
	"carpool-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UpdateUser(c *fiber.Ctx) error {
	var updateUser models.User

	// Parse and decode the JSON request body into the user struct
	if err := c.BodyParser(&updateUser); err != nil {
		log.Printf("Error parsing JSON request body: %v\n", err)
		return c.Status(400).JSON("Error parsing JSON: " + err.Error())
	}

	// validates user, display error if not validated
	if err := helpers.ValidateUser(updateUser); err != nil {
		log.Printf("Error validating user: %v\n", err)
		return c.Status(400).SendString("Error validating user")
	}

	// Update the user to the database, using id.
	result := database.Database.Db.Model(&models.User{}).Where("name = ?", updateUser.Name).Updates(&updateUser)

	// Check for errors during the database update
	if result.Error != nil {
		log.Printf("Error updating user: %v\n", result.Error)
		return c.Status(500).SendString("Error updating user")
	}

	// Check if any rows were affected during the update
	if result.RowsAffected == 0 {
		log.Printf("User with name %v not found\n", updateUser.Name)
		return c.Status(404).SendString("User not found")
	}

	// Log success and return the updated user as JSON.
	log.Printf("User with name %v successfully updated\n", updateUser.Name)
	return c.Status(200).JSON(helpers.CreateResponseUser(updateUser))
}
