package rides

import (
	"carpool-backend/database"
	"carpool-backend/helpers"
	"carpool-backend/models"
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Ride struct represents the ride model.
type Ride struct {
	ID            uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	HostUser      models.User `json:"host_user"`
	StartLocation string      `gorm:"size:255;not null;" json:"start_location"`
	EndLocation   string      `gorm:"size:255;not null;" json:"end_location"`
	StartTime     time.Time   `gorm:"not null;" json:"start_time"`
	TotalSeats    uint        `gorm:"not null;" json:"total_seats"`
	BookedSeats   uint        `gorm:"not null;" json:"booked_seats"`
	TotalPrice    uint        `gorm:"not null;" json:"total_price"`
}

// CreateResponseRide creates a response ride by mapping a ride model and user model.
func CreateResponseRide(rideModel models.Ride, user models.User) Ride {
	return Ride{
		ID:            rideModel.ID,
		HostUser:      user,
		StartLocation: rideModel.StartLocation,
		EndLocation:   rideModel.EndLocation,
		StartTime:     rideModel.StartTime,
		TotalSeats:    rideModel.TotalSeats,
		BookedSeats:   rideModel.BookedSeats,
		TotalPrice:    rideModel.TotalPrice,
	}
}

// FindRide finds a ride by its database ID.
func FindRide(id uint, ride *models.Ride) error {
	database.Database.Db.Find(&ride, "id = ?", id)
	if ride.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

// UpdateRide handles updating a ride based on the provided ID in the Fiber context.
func UpdateRide(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var ride models.Ride
	if err != nil {
		return c.Status(400).JSON("Please ensure that id is valid")
	}
	if err := FindRide(uint(id), &ride); err != nil {
		return c.Status(400).JSON("Error fetching the ride :" + err.Error())
	}
	type UpdateRide struct {
		HostUser      models.User `json:"host_user"`
		StartLocation string      `json:"start_location"`
		EndLocation   string      `json:"end_location"`
		TotalSeats    uint        `json:"total_seats"`
		BookedSeats   uint        `json:"booked_seats"`
		TotalPrice    uint        `json:"total_price"`
	}
	var updateData UpdateRide

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	// Update ride properties with new data.
	ride.HostUser = updateData.HostUser
	ride.StartLocation = updateData.StartLocation
	ride.EndLocation = updateData.EndLocation
	ride.TotalSeats = updateData.TotalSeats
	ride.BookedSeats = updateData.BookedSeats
	ride.TotalPrice = updateData.TotalPrice

	ride.ID = uint(id)

	// Save the updated ride to the database.
	updatedRide := database.Database.Db.Save(&ride)
	if updatedRide.Error != nil {
		return c.Status(500).SendString("Error updating ride")
	}

	// Create a response ride based on the updated ride and its host user.
	responseRide := CreateResponseRide(ride, ride.HostUser)

	// Convert the local response ride to a model.Ride for validation.
	modelsRide := ConvertToModelsRide(responseRide)

	// Validate the updated ride.
	if err := helpers.ValidateRide(modelsRide); err != nil {
		log.Printf("Error validating ride: %v\n", err)
		return c.Status(400).SendString("Error validating ride")
	}

	// Log success and return the updated ride as JSON.
	log.Printf("%v Your ride has been updated\n", ride.HostUser)
	return c.Status(200).JSON(responseRide)
}

// ConvertToModelsRide converts a local Ride struct to a model.Ride struct.
func ConvertToModelsRide(localRide Ride) models.Ride {
	return models.Ride{
		ID:            localRide.ID,
		HostUser:      localRide.HostUser,
		StartLocation: localRide.StartLocation,
		EndLocation:   localRide.EndLocation,
		StartTime:     localRide.StartTime,
		TotalSeats:    localRide.TotalSeats,
		BookedSeats:   localRide.BookedSeats,
		TotalPrice:    localRide.TotalPrice,
	}
}
