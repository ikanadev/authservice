package services

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gofiber/fiber/v2"
)

// WithAuth reads the Authorization token and set the ID in context
func WithAuth(roles []int8, JWTSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return NewErr(InvalidToken, "Access denied", fiber.StatusUnauthorized)
		}
		userID, role, err := CheckToken(JWTSecret, token)
		if err != nil {
			if err == jwt.ErrExpValidation {
				return NewErr(ExpiredToken, "JWT token expired", fiber.StatusUnauthorized)
			}
			return NewErr(InvalidToken, "Access denied", fiber.StatusUnauthorized)
		}
		for _, r := range roles {
			if r == role {
				c.Locals("userID", userID)
				c.Locals("role", role)
				return c.Next()
			}
		}
		return NewErr(InvalidToken, "Access denied", fiber.StatusUnauthorized)
	}
}
