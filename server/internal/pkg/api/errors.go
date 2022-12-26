package api

import "go/types"

const (
	InvalidRequestBodyErrorCode int = 1001
	UsernameRegisteredErrorCode     = 1002
)

type ErrorResponse[D any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *D     `json:"data"`
}

func InvalidRequestBody(message string) ErrorResponse[types.Nil] {
	return ErrorResponse[types.Nil]{
		Code:    InvalidRequestBodyErrorCode,
		Message: message,
		Data:    nil,
	}
}

func UsernameRegistered() ErrorResponse[types.Nil] {
	return ErrorResponse[types.Nil]{
		Code:    UsernameRegisteredErrorCode,
		Message: "username already registered",
		Data:    nil,
	}
}
