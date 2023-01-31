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

type GetHistoryRequestBody struct {
	WithUserID int64  `json:"withUserID" validate:"required"`
	BeforeKey  string `json:"beforeKey" validate:"omitempty,len=21"`
	Limit      int64  `json:"limit" validate:"required,min=0,max=500"`
}

type GetHistoryResponse struct {
	Messages []Message `json:"messages"`
	HasMore  bool      `json:"hasMore"`
}
