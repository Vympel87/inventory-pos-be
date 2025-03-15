package middleware

import (
	"time"

	"github.com/Vympel87/inventory-pos-be/internal/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiterMiddleware() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        100,
		Expiration: 60 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return exception.NewHttpException("Too Many Requests", 4299, fiber.StatusTooManyRequests, nil)
		},
	})
}
