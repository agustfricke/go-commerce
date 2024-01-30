package router

import (
	"github.com/agustfricke/go-commerce/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
  app.Get("/check", handlers.Check)
}
