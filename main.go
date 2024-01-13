package main

import (
	"carpool-backend/database"
	"carpool-backend/routes/bookings"
	"carpool-backend/routes/rides"
	"carpool-backend/routes/users"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// booking routes
	app.Post("/bookings", bookings.CreateBooking)
	app.Put("/bookings", bookings.EditBooking)
	app.Delete("/bookings", bookings.DeleteBooking)

	// ride routes
	app.Post("/rides", rides.CreateRide)
	app.Put("/rides", rides.UpdateRide)
	app.Patch("/rideStatus", rides.UpdateRideStatus)
	app.Delete("/rides", rides.DeleteRide)
	app.Get("/rides", rides.GetRides)
	app.Get("/search", rides.SearchRides)
	app.Get("/rides/:id", rides.GetRidesById)

	// user routes

	app.Put("/users", users.UpdateUser)
	app.Post("/users", users.CreateUser)
	app.Delete("/users/:id", users.DeleteUser)
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
