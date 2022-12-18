package auth

import (
	"errors"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/api_errors"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/connector"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/token"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

type registerRequestBody struct {
	Name     string `json:"name" validate:"required,max=60"`
	Username string `json:"username" validate:"required,max=20,username"`
	Password string `json:"password" validate:"required,max=72"`
}

type registerResponse struct {
	UserID int64  `json:"userID"`
	Token  string `json:"token"`
}

func Register(c *fiber.Ctx) error {
	var body registerRequestBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api_errors.InvalidRequestBody(err.Error()))
	}

	if err := validator.Validate.Struct(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api_errors.InvalidRequestBody(err.Error()))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	var userID int64
	err = connector.Pgx.
		QueryRow(c.Context(), `INSERT INTO "user" ("name", "username", "password") VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING "userID" `, body.Name, body.Username, hashedPassword).
		Scan(&userID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusConflict).JSON(api_errors.UsernameRegistered())
		}
		return err
	}

	authToken, err := token.CreateAuthToken(userID)
	if err != nil {
		return err
	}

	return c.JSON(registerResponse{
		UserID: userID,
		Token:  authToken,
	})
}
