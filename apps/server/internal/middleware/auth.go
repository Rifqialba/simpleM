package middleware

import (
	"strings"

	"github.com/Rifqialba/simplem/apps/server/internal/response"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return response.Error(
				c,
				fiber.StatusUnauthorized,
				"missing authorization header",
			)
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			},
		)

		if err != nil || !token.Valid {
			return response.Error(
				c,
				fiber.StatusUnauthorized,
				"invalid token",
			)
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Locals("user_id", claims["user_id"])

		return c.Next()
		
		
	}
	
}


