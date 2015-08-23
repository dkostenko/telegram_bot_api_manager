package tgbot

type Param struct {
	Key   string
	Value string
	Type  string
}

type ReplyMarkup struct {
	Keyboard        [][]string `json:"keyboard"`
	ResizeKeyboard  bool       `json:"resize_keyboard"`
	OneTimeKeyboard bool       `json:"one_time_keyboard"`
	Selective       bool       `json:"selective"`
	HideKeyboard    bool       `json:"hide_keyboard"`
	ForceReply      bool       `json:"force_reply"`
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
