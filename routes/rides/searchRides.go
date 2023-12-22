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

	const similarityThreshold = 0.4

	startLocation := c.Query("start_location")
	endLocation := c.Query("end_location")
	date := c.Query("date")
	//date in format DD-MM-YYYY

	var rides []models.Ride

	tx := database.Database.Db.Model(&models.Ride{})

	// First remove all rides that are full or have already started

	tx = tx.Where("booked_seats < total_seats AND start_time > NOW()")

	// Filtering by start location
	if startLocation != "" {
		tx = tx.Where("start_location ILIKE ? OR similarity(start_location, ?) > ?", "%"+startLocation+"%", startLocation, similarityThreshold)
	}

	// Filtering by end location
	if endLocation != "" {
		tx = tx.Where("end_location ILIKE ? OR similarity(end_location, ?) > ?", "%"+endLocation+"%", endLocation, similarityThreshold)
	}

	// Filtering by date
	if date != "" {
		inputDate, err := time.Parse("02-01-2006", date)

		if err != nil {
			log.Printf("Error parsing date: %v\n", err)
			return c.Status(400).SendString("Error parsing date")
		}

		timeLoc, err := time.LoadLocation("Asia/Calcutta")

		if err != nil {
			log.Printf("Error parsing date: %v\n", err)
			return c.Status(400).SendString("Error setting timezone")
		}

		inputDate = inputDate.In(timeLoc)
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
