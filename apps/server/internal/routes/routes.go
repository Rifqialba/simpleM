package routes

import (
	"github.com/Rifqialba/simplem/apps/server/internal/app"
	"github.com/Rifqialba/simplem/apps/server/internal/handler"
	"github.com/Rifqialba/simplem/apps/server/internal/middleware"
	"github.com/Rifqialba/simplem/apps/server/internal/response"
	"github.com/Rifqialba/simplem/apps/server/internal/user"

	"github.com/gofiber/fiber/v2"
)

func Register(
	appFiber *fiber.App,
	appContainer *app.App,
) {

	healthHandler := handler.New(appContainer)

	appFiber.Get("/health", healthHandler.Health)

	userRepo := user.NewRepository(appContainer.DB)

	userService := user.NewService(userRepo)

	userHandler := user.NewHandler(
		userService,
		appContainer.Config.JWTSecret,
	)

	appFiber.Post("/users", userHandler.Create)

	appFiber.Post("/login", userHandler.Login)

	appFiber.Get(
		"/me",
		middleware.Auth(appContainer.Config.JWTSecret),
		func(c *fiber.Ctx) error {

			return response.Success(
				c,
				fiber.StatusOK,
				fiber.Map{
					"user_id": c.Locals("user_id"),
				},
			)
		},
	)

	appFiber.Get("/panic", func(c *fiber.Ctx) error {
		panic("test panic")
	})
}