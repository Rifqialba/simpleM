package workspace

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

func (h *Handler) Create(
	c *fiber.Ctx,
) error {

	userID := c.Locals("user_id").(string)

	var req CreateWorkspaceRequest

	if err := c.BodyParser(&req); err != nil {

		return response.Error(
			c,
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	workspace, err := h.service.Create(
		c.Context(),
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

	return response.Success(
		c,
		fiber.StatusCreated,
		workspace,
	)
}