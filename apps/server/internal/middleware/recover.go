package middleware

import (
	"github.com/Rifqialba/simplem/apps/server/internal/response"
	"github.com/gofiber/fiber/v2"
)

func Recover() fiber.Handler {
	return func(c *fiber.Ctx) error {

		defer func() {
			if r := recover(); r != nil {

				_ = response.Error(
					c,
					fiber.StatusInternalServerError,
					"internal server error",
				)
			}
		}()

		return c.Next()
	}
}