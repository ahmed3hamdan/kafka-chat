package model

import (
	"context"
	"errors"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/db"
	"github.com/jackc/pgx/v5"
)

var UsernameRegisteredError = errors.New("username registered")
var UserNotFoundError = errors.New("user not found")

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

func GetUserById(ctx context.Context, userID int64) (User, error) {
	user := User{UserID: userID}
	err := db.Pgx.QueryRow(ctx, `SELECT "name", "username", "password" FROM "user" WHERE "userID" = $1`, userID).
		Scan(&user.Name, &user.Username, &user.Password)
	if err == pgx.ErrNoRows {
		return user, UserNotFoundError
	}
	return user, err
}

func GetUserByUsername(ctx context.Context, username string) (User, error) {
	user := User{Username: username}
	err := db.Pgx.QueryRow(ctx, `SELECT "userID", "name", "password" FROM "user" WHERE "username" = $1`, username).
		Scan(&user.UserID, &user.Name, &user.Password)
	if err == pgx.ErrNoRows {
		return user, UserNotFoundError
	}
	return user, err
}
