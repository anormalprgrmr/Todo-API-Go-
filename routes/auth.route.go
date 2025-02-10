package routes

import (
	"Todo-API/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupAuthRoutes initializes authentication routes
func SetupAuthRoutes(app *fiber.App) {
    app.Post("/register", handlers.Signup)
    app.Post("/login", handlers.Login)
}