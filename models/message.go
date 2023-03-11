package models

type Message struct {
	MessageID    int    `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	Message      string `json:"message" db:"message"`
	SentAt       string `json:"sent_at" db:"sent_at"`
	IsRegistered bool   `json:"is_registered" db:"is_registered"`
}

type MessageRes struct {
	LastMessageID int       `json:"last_id"`
	Messages      []Message `json:"messages"`
}

type GuestMessageReq struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

type MessagesFromLast struct {
	LastID int `form:"last_id"`
	Limit  int `form:"limit"`
}

type DeletedMessage struct {
	DeletedID int `json:"deleted_id"`
}

type UserMessageReq struct {
	Message string `json:"message"`
}
