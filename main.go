package main

import (
	"log"
	"net/http"

	"github.com/PrinceNarteh/go-users/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// loading environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to database
	configs.ConnectDB()

	// creating fiber instance
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(&fiber.Map{"message": "API running"})
	})

	app.Listen(":4000")
}
