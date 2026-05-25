package tab

import (
	"github.com/Rifqialba/simplem/apps/server/internal/realtime"
	"github.com/Rifqialba/simplem/apps/server/internal/response"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *Service

	realtimeManager *realtime.Manager
}

func NewHandler(
	service *Service,
	realtimeManager *realtime.Manager,
) *Handler {

	return &Handler{
		service: service,

		realtimeManager: realtimeManager,
	}
}

func (h *Handler) Create(
	c *fiber.Ctx,
) error {

	roomID := c.Params("roomId")

	userID := c.Locals("user_id").(string)

	var req CreateTabRequest

	if err := c.BodyParser(&req); err != nil {

		return response.Error(
			c,
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	tab, err := h.service.Create(
		c.Context(),
		roomID,
		userID,
		req,
	)

	if err != nil {

		return response.Error(
			c,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	h.realtimeManager.Broadcast(
		roomID,
		realtime.Event{
			Type: realtime.EventTabCreated,
			Payload: tab,
		},
	)

	return response.Success(
		c,
		fiber.StatusCreated,
		tab,
	)
}

func (h *Handler) ListByRoomID(
	c *fiber.Ctx,
) error {

	roomID := c.Params("roomId")

	tabs, err := h.service.ListByRoomID(
		c.Context(),
		roomID,
	)

	if err != nil {

		return response.Error(
			c,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	return response.Success(
		c,
		fiber.StatusOK,
		tabs,
	)
}

func (h *Handler) Activate(
	c *fiber.Ctx,
) error {

	roomID := c.Params("roomId")

	tabID := c.Params("tabId")

	err := h.service.ActivateTab(
		c.Context(),
		roomID,
		tabID,
	)

	if err != nil {

		return response.Error(
			c,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	h.realtimeManager.Broadcast(
		roomID,
		realtime.Event{
			Type: realtime.EventTabActivated,
			Payload: fiber.Map{
				"tab_id": tabID,
			},
		},
	)

	return response.Success(
		c,
		fiber.StatusOK,
		fiber.Map{
			"message": "tab activated",
		},
	)
}