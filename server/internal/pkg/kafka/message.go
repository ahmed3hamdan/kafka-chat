package kafka

type Message struct {
	Key        string `json:"key"`
	FromUserID int64  `json:"fromUserID"`
	ToUserID   int64  `json:"toUserID"`
	Content    string `json:"content"`
}
