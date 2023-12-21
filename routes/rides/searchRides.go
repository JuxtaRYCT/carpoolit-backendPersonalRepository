package rides

import (
	"carpool-backend/database"
	"carpool-backend/helpers"
	"carpool-backend/models"
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// This is the heart of the app - the search function.
// Full text search offered by Postgres for fuzzy search
// Exact matches for filtered fields
// Filters available: Date, Start Location, End Location
// TODO: Radial search for location, find a geocoding API to use

func SearchRides(c *fiber.Ctx) error {

	startLocation := c.Query("startLocation")
	endLocation := c.Query("endLocation")
	date := c.Query("date")
	//date in format DD-MM-YYYY

	var rides []models.Ride

	tx := database.Database.Db.Model(&models.Ride{})

	if startLocation != "" {
		tx = tx.Where("start_location ILIKE ?", "%"+startLocation+"%")
	}

	if endLocation != "" {
		tx = tx.Where("end_location ILIKE ?", "%"+endLocation+"%")
	}

	if date != "" {
		inputDate, err := time.Parse("02-01-2006", date)
		location := time.Now().Location()
		inputDate = inputDate.In(location)
		log.Printf("Date: %v\n", inputDate)
		if err != nil {
			return c.Status(400).SendString("Error parsing date")
		}

		startOfDay := time.Date(inputDate.Year(), inputDate.Month(), inputDate.Day(), 0, 0, 0, 0, inputDate.Location())
		endOfDay := time.Date(inputDate.Year(), inputDate.Month(), inputDate.Day(), 23, 59, 59, 0, inputDate.Location())

		tx = tx.Where("start_time BETWEEN ? AND ?", startOfDay, endOfDay)
	}

	err := tx.Find(&rides).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).SendString("No rides found")
	}

	if err != nil {
		return c.Status(500).SendString("Error fetching rides")
	}

	var responseRides []models.RideResponse

	for _, ride := range rides {
		responseRides = append(responseRides, helpers.CreateResponseRide(ride))
	}

	return c.Status(200).JSON(responseRides)
}
