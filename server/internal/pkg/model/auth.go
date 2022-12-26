package model

import (
	"context"
	"errors"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/db"
	"github.com/jackc/pgx/v5"
)

var UsernameRegisteredError = errors.New("username registered")

type User struct {
	UserID   int64
	Name     string
	Username string
	Password []byte
}

func InsertUser(ctx context.Context, user *User) error {
	err := db.Pgx.
		QueryRow(ctx, `INSERT INTO "user" ("name", "username", "password") VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING "userID" `, user.Name, user.Username, user.Password).
		Scan(&user.UserID)

	if err == pgx.ErrNoRows {
		return UsernameRegisteredError
	}

	return err
}
