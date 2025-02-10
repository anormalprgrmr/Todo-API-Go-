package middlewares

import (
	"Todo-API/database"
	"Todo-API/models"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthRequired(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	tokenStr := strings.Split(authHeader, "Bearer ")[1]
    secretKey := os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return []byte(secretKey), nil
    })

	if err != nil || !token.Valid {
        return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
    }

	claims := token.Claims.(jwt.MapClaims)
    var user models.User
    database.DataBase.First(&user, claims["user_id"])
    if user.ID == 0 {
        return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
    }

    c.Locals("userID", user.ID)
    return c.Next()

}