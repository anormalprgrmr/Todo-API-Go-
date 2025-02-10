package routes

import (
	"Todo-API/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
    app.Post("/signup", handlers.Signup)
    app.Post("/login", handlers.Login)
}