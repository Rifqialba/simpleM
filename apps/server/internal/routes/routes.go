package routes

import (
	"github.com/Rifqialba/simplem/apps/server/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/health", handler.Health)
}