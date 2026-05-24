package routes

import (
	"github.com/Rifqialba/simplem/apps/server/internal/app"
	"github.com/Rifqialba/simplem/apps/server/internal/handler"
	"github.com/Rifqialba/simplem/apps/server/internal/user"

	"github.com/gofiber/fiber/v2"
)

func Register(
	appFiber *fiber.App,
	appContainer *app.App,
) {
	h := handler.New(appContainer)

	appFiber.Get("/health", h.Health)

	userRepo := user.NewRepository(appContainer.DB)

	userService := user.NewService(userRepo)

	userHandler := user.NewHandler(userService)

	appFiber.Post("/users", userHandler.Create)
	
	appFiber.Get("/panic", func(c *fiber.Ctx) error {
	panic("test panic")
})
}