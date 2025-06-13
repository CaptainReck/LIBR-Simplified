package model

type Message struct {
	Id        string `json:"id"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
	Status    string `json:"status"`
}

type ModResponse struct {
	ModeratorId  int
	MessageId    string
	Status       string
	ResponseTime int
}

type MessageResponse struct {
	Id        string `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Status    string `json:"status"`
}
