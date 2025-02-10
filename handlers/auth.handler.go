package handlers

import (
	"Todo-API/database"
	"Todo-API/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string,error) {
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(password),10)
	return string(hashedPassword),err
}

func Signup(c *fiber.Ctx) error {
	var data map[string]string
	if err:=c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":"Failed!",
		})
	}

	hashedPassword, _ := HashPassword(data["password"])

	user := models.User{
		Username: data["username"],
		Password : hashedPassword,
	}

	database.DataBase.Create((&user))

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
    if err := c.BodyParser(&data); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
    }
	
	var user models.User
    database.DataBase.Where("username = ?", data["username"]).First(&user)

    if user.ID == 0 {
        return c.Status(404).JSON(fiber.Map{"error": "User not found"})
    }

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Incorrect password"})
    }

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	secretKey := os.Getenv("SECRET_KEY")
    t, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Could not generate token"})
    }

    return c.JSON(fiber.Map{"secretKey":os.Getenv("secretKey"),"token": t})
}