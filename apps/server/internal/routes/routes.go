package routes

import (
	"github.com/Rifqialba/simplem/apps/server/internal/app"
	"github.com/Rifqialba/simplem/apps/server/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func Register(
	appFiber *fiber.App,
	appContainer *app.App,
) {
	h := handler.New(appContainer)

	appFiber.Get("/health", h.Health)
}