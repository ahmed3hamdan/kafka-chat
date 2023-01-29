package kafka

import "time"

type Message struct {
	Key        string
	FromUserID int64
	ToUserID   int64
	Content    string
	CreatedAt  time.Time
}
