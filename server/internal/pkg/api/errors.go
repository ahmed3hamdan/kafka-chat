package api

import "go/types"

const (
	InvalidRequestBodyErrorKey = "invalid-request-body"
	UsernameRegisteredErrorKey = "username-registered"
	UserNotFoundErrorKey       = "user-not-found"
	PasswordMismatchErrorKey   = "password-mismatch"
	InvalidAuthTokenErrorKey   = "invalid-auth-token"
)

type ErrorResponse[D any] struct {
	Key     string `json:"key"`
	Message string `json:"message"`
	Data    *D     `json:"data"`
}

func InvalidRequestBody(message string) ErrorResponse[types.Nil] {
	return ErrorResponse[types.Nil]{
		Key:     InvalidRequestBodyErrorKey,
		Message: message,
		Data:    nil,
	}
}

func UsernameRegistered() ErrorResponse[types.Nil] {
	return ErrorResponse[types.Nil]{
		Key:     UsernameRegisteredErrorKey,
		Message: "username already registered",
		Data:    nil,
	}
}

func UserNotFound(message string) ErrorResponse[types.Nil] {
	return ErrorResponse[types.Nil]{
		Key:     UserNotFoundErrorKey,
		Message: message,
		Data:    nil,
	}
}

func PasswordMismatch() ErrorResponse[types.Nil] {
	return ErrorResponse[types.Nil]{
		Key:     PasswordMismatchErrorKey,
		Message: "hash and password mismatch",
		Data:    nil,
	}
}

func InvalidAuthToken(message string) ErrorResponse[types.Nil] {
	return ErrorResponse[types.Nil]{
		Key:     InvalidAuthTokenErrorKey,
		Message: message,
		Data:    nil,
	}
}
