package auth

import (
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/api"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/token"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func RequireAuthMiddleware(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")
	tokenString := strings.TrimPrefix(authorization, "Bearer ")

	userID, err := token.ValidateAuthToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(api.InvalidAuthToken(err.Error()))
	}

	c.Locals("userID", userID)

	return c.Next()
}
