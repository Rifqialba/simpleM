package middleware

import (
	"time"

	"github.com/Rifqialba/simplem/apps/server/internal/app"
	"github.com/gofiber/fiber/v2"
)

func Logger(appContainer *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {

		start := time.Now()

		err := c.Next()

		latency := time.Since(start)

		appContainer.Logger.Info().
			Str("request_id", c.Locals("request_id").(string)).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Int("status", c.Response().StatusCode()).
			Dur("latency", latency).
			Msg("incoming request")

		return err
	}
}