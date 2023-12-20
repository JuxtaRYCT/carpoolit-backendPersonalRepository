package main

import (
	"carpool-backend/database"
	"carpool-backend/routes/bookings"
	"carpool-backend/routes/rides"
	"carpool-backend/routes/users"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectToDB()

	SetupRoutes(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Fatal(app.Listen(":3000"))
}

func SetupRoutes(app *fiber.App) {
	// booking routes
	app.Put("/bookings/:id", bookings.EditBooking)
	app.Post("/bookings", bookings.CreateBooking)
	app.Post("/api/user", users.CreateUser)
	app.Post("/api/ride", rides.CreateRide)
}
