package main

import (
	"carpool-backend/database"
	routes "carpool-backend/routes/ride"
	"carpool-backend/routes/bookings"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// booking routes
	app.Delete("/bookings/:id", bookings.DeleteBooking)
	app.Post("/bookings", bookings.CreateBooking)
  app.Post("/createRide", routes.CreateRide)

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
