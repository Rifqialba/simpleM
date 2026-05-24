package response

import "github.com/gofiber/fiber/v2"

func Success(
	c *fiber.Ctx,
	status int,
	data interface{},
) error {

	return c.Status(status).JSON(
		fiber.Map{
			"success": true,
			"data":    data,
		},
	)
}

func Error(
	c *fiber.Ctx,
	status int,
	message string,
) error {

	return c.Status(status).JSON(
		fiber.Map{
			"success": false,
			"error": fiber.Map{
				"message": message,
			},
		},
	)
}