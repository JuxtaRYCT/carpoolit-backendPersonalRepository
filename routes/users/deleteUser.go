package users

import (
	"carpool-backend/database"
	"carpool-backend/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Delete a user based on the id as a query line parameter (empty body)

func DeleteUser(c *fiber.Ctx) error {
	//Raise an error if wrong request type
	if c.Method() != "DELETE" {
		return c.Status(400).SendString("Invalid request type")
	}

	//Get the user id from the URL parameter
	userId, err := strconv.Atoi(c.Params("id")) // Get the user id from the URL path and convert to integer

	if err != nil {
		log.Printf("Error converting user id to integer: %v\n", err)
		return c.Status(400).SendString("Invalid user id")
	} //to check if there's any error in converting the user id to integer

	//Delete the user from the database
	result := database.Database.Db.Delete(&models.User{}, userId)

	if result.Error != nil {
		log.Printf("Error deleting user: %v\n", result.Error)
		return c.Status(500).SendString("Could not delete user - Internal Server Error")
	} //If the user isn't deleted properly the function will return an error, and we'll know there was a problem during deletion

	if result.RowsAffected == 0 {
		log.Printf("Error deleting user: user not found - %v\n", userId)
		return c.Status(404).SendString("user not found")
	} //If none of the rows change ,then there was no user to be deleted, and we'll return a 404 status code

	log.Printf("User with id %v deleted", userId)
	return c.Status(200).SendString("User deleted") // in case it is successfully deleted we'll print user deleted
}
