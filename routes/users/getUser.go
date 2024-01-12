package users

import (
	"carpool-backend/database"
	"carpool-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	// Get user ID from the request parameters
	userID := c.Params("id")

	var user models.User

	// Fetch the user by ID along with preloading the past rides
	if err := database.Database.Db.Preload("Rides").First(&user, userID).Error; err != nil {
		log.Printf("Error finding user: %v\n", err)
		return c.Status(404).SendString("User not found")
	}

	// Return the user as JSON
	return c.Status(200).JSON(user)
}
