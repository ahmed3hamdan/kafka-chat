package auth

import (
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/api"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/model"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/token"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var body api.LoginRequestBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api.InvalidRequestBody(err.Error()))
	}

	if err := validator.Validate.Struct(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api.InvalidRequestBody(err.Error()))
	}

	user, err := model.GetUserByUsername(c.Context(), body.Username)
	if err == model.UserNotFoundError {
		return c.Status(fiber.StatusNotFound).JSON(api.NotFound(err.Error()))
	} else if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(body.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return c.Status(fiber.StatusUnauthorized).JSON(api.PasswordMismatch())
	} else if err != nil {
		return err
	}

	var authToken string
	authToken, err = token.CreateAuthToken(user.UserID)
	if err != nil {
		return err
	}

	return c.JSON(api.AuthResponse{
		UserID: user.UserID,
		Token:  authToken,
	})
}
