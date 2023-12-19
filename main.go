package main

import (
	"carpool-backend/database"
	"carpool-backend/routes/rides"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectToDB()
	setUpRoutes(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	log.Fatal(app.Listen(":3000"))
}

func setUpRoutes(app *fiber.App) {
	app.Put("/api/ride/:id", rides.UpdateRide)
}
