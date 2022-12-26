package auth

import (
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/api"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func GetSelfInfo(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int64)
	user, err := model.GetUserById(c.Context(), userID)
	if err == model.UserNotFoundError {
		return c.Status(fiber.StatusNotFound).JSON(api.NotFound(err.Error()))
	} else if err != nil {
		return err
	}
	return c.JSON(api.GetSelfInfoResponse{
		UserID:   user.UserID,
		Name:     user.Name,
		Username: user.Username,
	})
}
