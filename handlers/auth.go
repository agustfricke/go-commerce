package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/agustfricke/go-commerce/database"
	"github.com/agustfricke/go-commerce/models"
	"github.com/agustfricke/go-commerce/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt"
)

func Check(c *fiber.Ctx) error {
  return c.SendString("Hellooo!")
}

func SignUp(c *fiber.Ctx) error {
	var payload *models.SignUpInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	newUser, err := database.AddUser(models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: string(hashedPassword),
	})

  if err != nil {
      log.Fatal(err)
    }

	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)
  expDuration := time.Hour * 24

  claims["sub"] = newUser.ID
  claims["exp"] = now.Add(expDuration).Unix()
  claims["iat"] = now.Unix()
  claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(utils.GetEnv("SECRET_KEY")))

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("generating JWT Token failed: %v", err)})
	}

  // SendEmail(tokenString, payload.Email)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"token": tokenString, "data": newUser})
}
