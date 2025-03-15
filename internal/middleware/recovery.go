package middleware

import (
	"github.com/Vympel87/inventory-pos-be/config"
	"github.com/Vympel87/inventory-pos-be/internal/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// ErrorMiddleware menangani semua error dan mengembalikan response JSON sesuai jenis error
func ErrorMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		// Jika error tidak nil, lakukan penanganan error
		if err != nil {
			// Mengabaikan error terkait favicon.ico
			if c.Path() == "/favicon.ico" {
				return nil
			}

			// Handle HttpException
			if httpErr, ok := err.(*exception.HttpException); ok {
				logError(httpErr) // Logging
				return c.Status(httpErr.StatusCode).JSON(fiber.Map{
					"status":    "error",
					"errorCode": httpErr.ErrorCode,
					"message":   httpErr.Message,
				})
			}

			// Handle general error (seperti internal server error)
			logError(err)
			return c.Status(500).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		return nil
	}
}

// logError untuk log error dengan level yang sesuai
func logError(err interface{}) {
	switch e := err.(type) {
	case *exception.HttpException:
		// Log error dengan informasi lengkap
		entry := config.Log.WithFields(logrus.Fields{
			"errorCode": e.ErrorCode,
			"message":   e.Message,
			"errors":    e.Errors,
		})
		if e.StatusCode >= 500 {
			entry.Error("An internal server error occurred")
		} else {
			entry.Warn("Client error occurred")
		}
	default:
		// Log error umum
		config.Log.WithField("error", e).Error("An error occurred")
	}
}
