package main

import (
	"log"

	"github.com/PrinceNarteh/go-users/configs"
	"github.com/PrinceNarteh/go-users/routes"
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

	// routes
	routes.UserRoute(app)

	app.Listen(":4000")
}
