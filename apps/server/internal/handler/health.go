package handler

import (
	"context"

	"github.com/Rifqialba/simplem/apps/server/internal/app"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	App *app.App
}

func New(app *app.App) *Handler {
	return &Handler{
		App: app,
	}
}

func (h *Handler) Health(c *fiber.Ctx) error {
	ctx := context.Background()

	dbStatus := "ok"
	redisStatus := "ok"

	if err := h.App.DB.Ping(ctx); err != nil {
		dbStatus = "error"
	}

	if err := h.App.Redis.Ping(ctx).Err(); err != nil {
		redisStatus = "error"
	}

	return c.JSON(fiber.Map{
		"status": "ok",

		"services": fiber.Map{
			"database": dbStatus,
			"redis":    redisStatus,
		},
	})
}