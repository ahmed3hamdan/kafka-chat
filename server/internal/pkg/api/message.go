package api

type SendMessageRequestBody struct {
	UserID  int64  `json:"userID" validate:"required"`
	Content string `json:"content" validate:"required"`
}
