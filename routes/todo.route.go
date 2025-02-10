package routes

import (
	"Todo-API/handlers"
	"Todo-API/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupTodoRoutes(app fiber.Router) {

    app.Post("/", middlewares.AuthRequired,handlers.CreateTodo)
    app.Get("/",middlewares.AuthRequired, handlers.GetTodos)
    // todoGroup.Get("/:id", handlers.GetTodoByID)
    // todoGroup.Put("/:id", handlers.UpdateTodo)
    // todoGroup.Delete("/:id", handlers.DeleteTodo)
}
