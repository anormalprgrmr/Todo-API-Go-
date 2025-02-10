package main

import (
	"Todo-API/database"
	"Todo-API/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	database.ConnectDB()
	app := fiber.New()

	app.Get("/hello",func (c *fiber.Ctx) error {
		return c.SendString("Hello ! this is test")
	})
	routes.SetupAuthRoutes(app)
	routes.SetupTodoRoutes(app)

	app.Listen(":3100")
}