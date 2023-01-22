package kafka

type Message struct {
	SendID    int64  `json:"sendID"`
	ReceiveID int64  `json:"receiveID"`
	Content   string `json:"content"`
}
