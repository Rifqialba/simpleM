package routes

import (
	"github.com/Rifqialba/simplem/apps/server/internal/app"
	"github.com/Rifqialba/simplem/apps/server/internal/handler"
	"github.com/Rifqialba/simplem/apps/server/internal/middleware"
	"github.com/Rifqialba/simplem/apps/server/internal/realtime"
	"github.com/Rifqialba/simplem/apps/server/internal/response"
	"github.com/Rifqialba/simplem/apps/server/internal/user"
	"github.com/Rifqialba/simplem/apps/server/internal/tab"
	"github.com/Rifqialba/simplem/apps/server/internal/workspace"
	"github.com/Rifqialba/simplem/apps/server/internal/room"
	ws "github.com/gofiber/contrib/websocket"
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

	tabRepo := tab.NewRepository(appContainer.DB)

	tabService := tab.NewService(tabRepo)

	tabHandler := tab.NewHandler(
	tabService,
	appContainer.RealtimeManager,
)


	workspaceRepo := workspace.NewRepository(appContainer.DB)

	workspaceService := workspace.NewService(workspaceRepo)

	workspaceHandler := workspace.NewHandler(workspaceService)

	roomRepo := room.NewRepository(appContainer.DB)

	roomService := room.NewService(roomRepo)

	roomHandler := room.NewHandler(roomService)

	realtimeHandler := realtime.NewHandler(
	appContainer.RealtimeManager,
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

	appFiber.Get(
		"/ws/:roomId",
		ws.New(realtimeHandler.Handle),
	)

	appFiber.Get("/panic", func(c *fiber.Ctx) error {
		panic("test panic")
	})

	appFiber.Post(
	"/rooms/:roomId/tabs",
	middleware.Auth(appContainer.Config.JWTSecret),
	tabHandler.Create,
	)

	appFiber.Get(
	"/rooms/:roomId/tabs",
	middleware.Auth(appContainer.Config.JWTSecret),
	tabHandler.ListByRoomID,
	)

	appFiber.Put(
	"/rooms/:roomId/tabs/:tabId/activate",
	middleware.Auth(appContainer.Config.JWTSecret),
	tabHandler.Activate,
	)

	appFiber.Post(
	"/workspaces",
	middleware.Auth(appContainer.Config.JWTSecret),
	workspaceHandler.Create,
	)

    appFiber.Post(
	"/workspaces/:workspaceId/rooms",
	middleware.Auth(appContainer.Config.JWTSecret),
	roomHandler.Create,
	)
}