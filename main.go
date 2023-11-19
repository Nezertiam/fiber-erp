package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("SERVER_PORT")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Port: " + port)
	})

	app.Listen(":8000")
}
