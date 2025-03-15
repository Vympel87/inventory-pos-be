package middleware

import (
	"github.com/Vympel87/inventory-pos-be/config"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SuccessOperation() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err == nil {
			config.Log.WithFields(logrus.Fields{
				"method": c.Method(),
				"path":   c.Path(),
				"status": c.Response().StatusCode(),
			}).Info("Successful request")
		}
		return err
	}
}
