package main

import (
	"carpool-backend/database"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	database.ConnectToDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Fatal(app.Listen(":3000"))
}
