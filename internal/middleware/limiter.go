package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)


func Login_Limitter()  fiber.Handler {
			login_limiter := limiter.New(limiter.Config{
				Max: 5,
				Expiration: 1 * time.Minute,
				LimitReached:  func (c *fiber.Ctx) error {
					return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
						"error" : "Too many resquest for login Attempt.please try later",
					})
				},
			})
	return login_limiter
}

