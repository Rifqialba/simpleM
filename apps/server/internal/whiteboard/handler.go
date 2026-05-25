package whiteboard

import (
	"github.com/Rifqialba/simplem/apps/server/internal/response"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *Service
}

func NewHandler(
	service *Service,
) *Handler {

	return &Handler{
		service: service,
	}
}

func (h *Handler) Save(
	c *fiber.Ctx,
) error {

	tabID := c.Params("tabId")

	var req SaveWhiteboardRequest

	if err := c.BodyParser(&req); err != nil {

		return response.Error(
			c,
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	whiteboard, err := h.service.Save(
		c.Context(),
		tabID,
		req,
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
		whiteboard,
	)
}

func (h *Handler) FindByTabID(
	c *fiber.Ctx,
) error {

	tabID := c.Params("tabId")

	whiteboard, err := h.service.FindByTabID(
		c.Context(),
		tabID,
	)

	if err != nil {

		return response.Error(
			c,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	return response.Success(
		c,
		fiber.StatusOK,
		whiteboard,
	)
}