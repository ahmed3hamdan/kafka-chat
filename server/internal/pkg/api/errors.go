package api

import "go/types"

const (
	InvalidRequestBodyErrorCode = 1001
	UsernameRegisteredErrorCode = 1002
	InvalidParamsErrorCode      = 1003
	NotFoundErrorCode           = 1004
	PasswordMismatchErrorCode   = 1005
	InvalidAuthTokenErrorCode   = 1006
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

func InvalidRequestParams(message string) ErrorResponse[types.Nil] {
	return ErrorResponse[types.Nil]{
		Code:    InvalidParamsErrorCode,
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

func NotFound(message string) ErrorResponse[types.Nil] {
	return ErrorResponse[types.Nil]{
		Code:    NotFoundErrorCode,
		Message: message,
		Data:    nil,
	}
}

func PasswordMismatch() ErrorResponse[types.Nil] {
	return ErrorResponse[types.Nil]{
		Code:    PasswordMismatchErrorCode,
		Message: "hash and password mismatch",
		Data:    nil,
	}
}

func InvalidAuthToken(message string) ErrorResponse[types.Nil] {
	return ErrorResponse[types.Nil]{
		Code:    InvalidAuthTokenErrorCode,
		Message: message,
		Data:    nil,
	}
}
