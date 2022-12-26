package token

import (
	"errors"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/config"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
)

var InvalidTokenSigningMethodError = errors.New("invalid token signing method")
var InvalidTokenError = errors.New("invalid token")
var InvalidClaimsError = errors.New("invalid claims")

var jwtSecret = []byte(config.JwtSecret)

func CreateAuthToken(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud": "auth",
		"sub": strconv.FormatInt(userID, 10),
	})
	return token.SignedString(jwtSecret)
}

func ValidateAuthToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, InvalidTokenSigningMethodError
		}
		return jwtSecret, nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["aud"] != "auth" {
			return 0, InvalidClaimsError
		}
		userID, err := strconv.ParseInt(claims["sub"].(string), 10, 64)
		if err != nil {
			return 0, InvalidClaimsError
		}
		return userID, nil
	}
	return 0, InvalidTokenError
}
