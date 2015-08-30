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
	MessageId           int              `json:message_id`
	Text                string           `json:"text"`
	From                *User            `json:"from"`
	Date                int              `json:"date"`
	Chat                *UserOrGroupChat `json:"chat"`
	ForwardFrom         *User            `json:"forward_from"`
	ForwardDate         int              `json:"forward_date"`
	ReplyToMessage      *Message         `json:"reply_to_message"`
	NewChatTitle        string           `json:"new_chat_title"`
	Voice               *Voice           `json:"voice"`
	Location            *Location        `json:"location"`
	Photo               []*PhotoSize     `json:"photo"`
	Audio               *Audio           `json:"audio"`
	Document            *Document        `json:"document"`
	Sticker             *Sticker         `json:"sticker"`
	Video               *Video           `json:"video"`
	Caption             string           `json:"caption"`
	Contact             *Contact         `json:"contact"`
	NewChatParticipant  *User            `json:"new_chat_participant"`
	LeftChatParticipant *User            `json:"left_chat_participant"`
	NewChatPhoto        []*PhotoSize     `json:"new_chat_photo"`
	DeleteChatPhoto     bool             `json:"delete_chat_photo"`
	GroupChatCreated    bool             `json:"group_chat_created"`
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

type UserProfilePhotos struct {
	TotalCount int            `json:"total_count"`
	Photos     [][]*PhotoSize `json:"photos"`
}

type PhotoSize struct {
	FileId   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}

type Voice struct {
	Duration  int    `json:"duration"`
	Mime_type string `json:"mime_type"`
	File_id   string `json:"file_id"`
	File_size int    `json:"file_size"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Sticker struct {
	FileId   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Thumb    *PhotoSize `json:"thumb"`
	FileSize int        `json:"file_size"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserId      int    `json:"user_id"`
}

type Video struct {
	FileId   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb"`
	MimeType string     `json:"mime_type"`
	FileSize int        `json:"file_size"`
}

type Document struct {
	FileId   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb"`
	FileName string     `json:"file_name"`
	MimeType string     `json:"mime_type"`
	FileSize int        `json:"file_size"`
}

type Audio struct {
	FileId    string `json:"file_id"`
	Duration  string `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
	MimeType  string `json:"mime_type"`
	FileSize  int    `json:"file_size"`
}

type Update struct {
	UpdateId string   `json:"update_id"`
	Message  *Message `json:"message"`
}
