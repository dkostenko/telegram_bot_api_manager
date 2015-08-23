// Package telegram_bot_api provides methods for communicating with Telegram Bot API.
package tgbot

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type (
	ForwardMessageRequest struct {
		ChatId     int
		FromChatId int
		MessageId  int
	}

	ForwardMessageResponse struct {
		Ok          bool    `json:"ok"`
		Result      Message `json:"result"`
		ErrorCode   int     `json:"error_code"`
		Description string  `json:"description"`
	}
)

func (this *ForwardMessageRequest) Send(secret string) (*ForwardMessageResponse, error) {
	url := getUrl("forwardMessage", secret)

	params, err := this.getParams()
	if err != nil {
		return nil, err
	}

	respJson, err := sendPost(url, params)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(strings.NewReader(respJson))

	var resp *ForwardMessageResponse
	if err := dec.Decode(&resp); err == io.EOF {
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *ForwardMessageRequest) getParams() ([]*Param, error) {
	res := []*Param{
		{Type: "string", Key: "chat_id", Value: strconv.Itoa(this.ChatId)},
		{Type: "string", Key: "from_chat_id", Value: strconv.Itoa(this.FromChatId)},
		{Type: "string", Key: "message_id", Value: strconv.Itoa(this.MessageId)},
	}

	return res, nil
}
