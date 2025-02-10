package routes

import (
	"Todo-API/handlers"
	"Todo-API/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupTodoRoutes(app *fiber.App) {
    todoGroup := app.Group("/todos", middlewares.AuthRequired)

    todoGroup.Post("/", handlers.CreateTodo)
    todoGroup.Get("/", handlers.GetTodos)
    // todoGroup.Get("/:id", handlers.GetTodoByID)
    // todoGroup.Put("/:id", handlers.UpdateTodo)
    // todoGroup.Delete("/:id", handlers.DeleteTodo)
}
