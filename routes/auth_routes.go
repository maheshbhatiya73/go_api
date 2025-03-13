package routes

import (
	"goapi/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/api/auth")
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)
}
