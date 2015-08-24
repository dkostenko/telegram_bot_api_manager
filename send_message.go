package tgbot

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type (
	SendMessageRequest struct {
		ChatId           int
		Text             string
		ReplyToMessageId int
		ReplyMarkup
	}

	SendMessageResponse struct {
		Ok          bool    `json:"ok"`
		Result      Message `json:"result"`
		ErrorCode   int     `json:"error_code"`
		Description string  `json:"description"`
	}
)

func (this *SendMessageRequest) Send(secret string) (*SendMessageResponse, error) {
	url := getUrl("sendMessage", secret)

	params, err := this.getParams()
	if err != nil {
		return nil, err
	}

	respJson, err := sendPost(url, params)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(strings.NewReader(respJson))

	var resp *SendMessageResponse
	if err := dec.Decode(&resp); err == io.EOF {
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *SendMessageRequest) getParams() ([]*Param, error) {
	replyMarkupJson, _ := json.Marshal(this.ReplyMarkup)

	res := []*Param{
		{Type: "string", Key: "chat_id", Value: strconv.Itoa(this.ChatId)},
		{Type: "string", Key: "text", Value: this.Text},
		{Type: "string", Key: "reply_markup", Value: string(replyMarkupJson)},
		{Type: "string", Key: "reply_to_message_id", Value: string(this.ReplyToMessageId)},
	}

	return res, nil
}
