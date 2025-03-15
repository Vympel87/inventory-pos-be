package app

import (
	"fmt"
	"os"

	"github.com/Vympel87/inventory-pos-be/internal/exception"
	"github.com/Vympel87/inventory-pos-be/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()

	// Middleware
	app.Use(middleware.CORSMiddleware())
	app.Use(middleware.RateLimiterMiddleware())
	app.Use(middleware.ErrorMiddleware())
	app.Use(middleware.SuccessOperation())

	// Dummy route
	app.Get("/hello_world", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	app.Get("/error", func(c *fiber.Ctx) error {
		return exception.NewHttpException("Oops!", 4001, 400, nil)
	})

	// Run app
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
