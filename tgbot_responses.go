// Package telegram_bot_api provides methods for communicating with Telegram Bot API.
package tgbot

type SendMessageResponse struct {
	Ok          bool    `json:"ok"`
	Result      Message `json:"result"`
	ErrorCode   int     `json:"error_code"`
	Description string  `json:"description"`
}

type Message struct {
	MessageId      int              `json:message_id`
	Text           string           `json:"text"`
	User           *User            `json:"from"`
	Date           int              `json:"date"`
	Chat           *UserOrGroupChat `json:"chat"`
	ForwardFrom    *User            `json:"forward_from"`
	ForwardDate    int              `json:"forward_date"`
	ReplyToMessage *Message         `json:"reply_to_message"`
	NewChatTitle   string           `json:"new_chat_title"`
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type UserOrGroupChat struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Title     string `json:"title"`
}
