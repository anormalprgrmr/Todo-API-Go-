package handlers

import (
	"Todo-API/database"
	"Todo-API/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
    }

	UserID := c.Locals("userID").(uint)
	todo.UserID = UserID

	database.DataBase.Create(&todo)
	return c.JSON(todo)
}

func GetTodos(c *fiber.Ctx) error {
    userID := c.Locals("userID").(uint)
    var todos []models.Todo

    database.DataBase.Where("user_id = ?", userID).Find(&todos)
    return c.JSON(todos)
}