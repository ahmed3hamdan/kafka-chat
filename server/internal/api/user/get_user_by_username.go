package user

import (
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/api"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/model"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func GetUserByUsername(c *fiber.Ctx) error {
	var body api.GetUserByUsernameRequestBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api.InvalidRequestBody(err.Error()))
	}

	if err := validator.Validate.Struct(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api.InvalidRequestBody(err.Error()))
	}

	user, err := model.GetUserByUsername(c.Context(), body.Username)
	if err == model.UserNotFoundError {
		return c.Status(fiber.StatusNotFound).JSON(api.UserNotFound(err.Error()))
	} else if err != nil {
		return err
	}

	return c.JSON(api.GetUserByUsernameResponse{
		UserID:   user.UserID,
		Name:     user.Name,
		Username: user.Username,
	})
}
