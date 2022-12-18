package api_errors

import "go/types"

const (
	InvalidRequestBodyCode int = 1001
	UsernameRegisteredCode     = 1002
)

type ResponseError[D any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *D     `json:"data"`
}

func InvalidRequestBody(message string) ResponseError[types.Nil] {
	return ResponseError[types.Nil]{
		Code:    InvalidRequestBodyCode,
		Message: message,
		Data:    nil,
	}
}

func UsernameRegistered() ResponseError[types.Nil] {
	return ResponseError[types.Nil]{
		Code:    UsernameRegisteredCode,
		Message: "username already registered",
		Data:    nil,
	}
}
