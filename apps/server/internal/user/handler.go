package user

import (
	"github.com/Rifqialba/simplem/apps/server/internal/response"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req CreateUserRequest

	if err := c.BodyParser(&req); err != nil {

		return response.Error(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
		)
	}

	user, err := h.service.Create(
		c.Context(),
		req,
	)

	if err != nil {

		return response.Error(
			c,
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	resp := UserResponse{
		ID:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		AvatarURL:   user.AvatarURL,
		IsActive:    user.IsActive,
	}

	return response.Success(
		c,
		fiber.StatusCreated,
		resp,
	)
}