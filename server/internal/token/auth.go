package token

import (
	"github.com/ahmed3hamdan/kafka-chat/server/internal/config"
	"github.com/golang-jwt/jwt/v4"
)

func CreateAuthToken(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud": "auth",
		"sub": userID,
	})
	return token.SignedString([]byte(config.JwtSecret))
}
