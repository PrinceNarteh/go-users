package routes

import (
	"github.com/PrinceNarteh/go-users/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Get("/:userId", controllers.GetUser)
}
