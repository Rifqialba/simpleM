package server

import (
	"github.com/Rifqialba/simplem/apps/server/internal/app"
	"github.com/Rifqialba/simplem/apps/server/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func New(
	appContainer *app.App,
) *fiber.App {

	app := fiber.New()

	app.Use(middleware.Recover())

	app.Use(middleware.RequestID())

	app.Use(middleware.Logger(appContainer))

	return app
}