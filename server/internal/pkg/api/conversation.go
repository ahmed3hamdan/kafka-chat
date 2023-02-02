package api

import "time"

type Conversation struct {
	Key                   string    `json:"key"`
	WithUserID            int64     `json:"withUserID"`
	WithName              string    `json:"withName"`
	WithUsername          string    `json:"withUsername"`
	LastMessageFromUserID int64     `json:"lastMessageFromUserID"`
	LastMessageContent    string    `json:"lastMessageContent"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}

type ListConversationsRequestBody struct {
	PageKey string `json:"pageKey" validate:"omitempty,len=21"`
	Limit   int    `json:"limit" validate:"required,min=0,max=500"`
}

type ListConversationsResponse struct {
	Entries     []Conversation `json:"entries"`
	NextPageKey *string        `json:"nextPageKey"`
}
