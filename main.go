package main

import (
	"carpool-backend/database"
	routes "carpool-backend/routes/ride"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectToDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	//Endpoint to create a ride
	app.Post("/createRide", routes.CreateRide)

	log.Fatal(app.Listen(":3000"))
}
