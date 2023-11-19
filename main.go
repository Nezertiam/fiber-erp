package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("No .env file found")
	}

	server_port := os.Getenv("SERVER_PORT")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Port: " + server_port)
	})

	app.Listen(":" + server_port)
}
