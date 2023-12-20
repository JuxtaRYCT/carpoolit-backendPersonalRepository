package main

import (
	"carpool-backend/database"

	"carpool-backend/routes/rides"
	"carpool-backend/routes/bookings"

	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// booking routes
	app.Delete("/bookings/:id", bookings.DeleteBooking)
	app.Post("/bookings", bookings.CreateBooking)

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

func setUpRoutes(app *fiber.App) {
	app.Put("/api/ride/:id", rides.UpdateRide)
}
