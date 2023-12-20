package rides

import (
	"carpool-backend/database"
	"carpool-backend/models"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// DeleteRide handles the deletion of a ride from the database
// Accepts a ride id to be passed as a path parameter in the URL
// Returns an HTTP status code and a message indicating the result of the operation
//
// Parameters:
//   - c: Fiber context
//
// Returns:
//   - error: An error message indicating the result of the operation and an HTTP status code
func DeleteRide(c *fiber.Ctx) error {
	rideId, err := strconv.Atoi(c.Params("id")) // Get the ride id from the URL path

	if err != nil {
		log.Printf("Error converting ride id to integer: %v\n", err)
		return c.Status(400).SendString("Invalid ride id")
	} //to check if there's any error in converting the ride id to integer

	result := database.Database.Db.Delete(&models.Ride{}, rideId) // Delete the ride from the database

	if result.Error != nil {
		log.Printf("Error deleting ride: %v\n", result.Error)
		return c.Status(500).SendString("Could not delete ride - Internal Server Error")
	} //If the ride isn't deleted properly the function will return an error, and we'll know there was a problem during deletion

	if result.RowsAffected == 0 {
		log.Printf("Error deleting ride: Ride not found - %v\n", rideId)
		return c.Status(404).SendString("Ride not found")
	} //If none of the rows change ,then there was no ride to be deleted, and we'll return a 404 status code

	responseMessage := fmt.Sprintf("Ride with id %v deleted", rideId)

	log.Println(responseMessage)
	return c.Status(200).SendString(fmt.Sprintf("Ride with id %v deleted", rideId)) // in case it is successfully deleted we'll print ride deleted and the ride id

}
