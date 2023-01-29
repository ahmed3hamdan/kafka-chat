package api

import "time"

type Message struct {
	Key        string    `json:"key"`
	FromUserID int64     `json:"fromUserID"`
	ToUserID   int64     `json:"toUserID"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
}

type SendMessageRequestBody struct {
	UserID  int64  `json:"userID" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type SendMessageResponse struct {
	MessageKey string `json:"messageKey"`
}
