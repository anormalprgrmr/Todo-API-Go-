package main

import (
	"Todo-API/database"
	"Todo-API/routes"
	"log"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "3100"
	}

	log.Printf("🚀 Server is running on http://localhost:%s", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}

}