package auth

import (
	"errors"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/api"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/model"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/token"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var body api.RegisterRequestBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.InvalidRequestBody(err.Error()))
	}

	if err := validator.Validate.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.InvalidRequestBody(err.Error()))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{
		Name:     body.Name,
		Username: body.Username,
		Password: hashedPassword,
	}

	if err = model.InsertUser(c.Context(), &user); err != nil {
		if errors.Is(err, model.UsernameRegisteredError) {
			return c.Status(fiber.StatusBadRequest).JSON(api.UsernameRegistered())
		}
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
