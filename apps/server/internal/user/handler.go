package user

import (
	"github.com/Rifqialba/simplem/apps/server/internal/response"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service    *Service
	jwtSecret string
}

func NewHandler(
	service *Service,
	jwtSecret string,
) *Handler {

	return &Handler{
		service:    service,
		jwtSecret: jwtSecret,
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

func (h *Handler) Login(c *fiber.Ctx) error {

	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {

		return response.Error(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
		)
	}

	resp, err := h.service.Login(
		c.Context(),
		req,
		h.jwtSecret,
	)

	if err != nil {

		return response.Error(
			c,
			fiber.StatusUnauthorized,
			"invalid credentials",
		)
	}

	return response.Success(
		c,
		fiber.StatusOK,
		resp,
	)
}