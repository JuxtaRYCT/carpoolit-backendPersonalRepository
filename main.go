package main

import (
	"carpool-backend/database"

	"carpool-backend/routes/bookings"
	"carpool-backend/routes/rides"

	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// booking routes
	app.Delete("/bookings/:id", bookings.DeleteBooking)
	app.Post("/bookings", bookings.CreateBooking)
	app.Put("/api/rides/:id", rides.UpdateRide)
	app.Put("/bookings/:id", bookings.EditBooking)
}

func main() {
	app := fiber.New()

	database.ConnectToDB()

	SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	log.Fatal(app.Listen(":3000"))
}
