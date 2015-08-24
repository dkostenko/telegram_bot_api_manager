package tgbot

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type (
	SendChatActionRequest struct {
		ChatId int
		Action string
	}

	SendChatActionResponse struct {
		Ok          bool   `json:"ok"`
		Result      bool   `json:"result"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	}
)

func (this *SendChatActionRequest) Send(secret string) (*SendChatActionResponse, error) {
	url := getUrl("sendChatAction", secret)

	params, err := this.getParams()
	if err != nil {
		return nil, err
	}

	respJson, err := sendPost(url, params)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(strings.NewReader(respJson))

	var resp *SendChatActionResponse
	if err := dec.Decode(&resp); err == io.EOF {
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *SendChatActionRequest) getParams() ([]*Param, error) {
	res := []*Param{
		{Type: "string", Key: "chat_id", Value: strconv.Itoa(this.ChatId)},
		{Type: "string", Key: "action", Value: this.Action},
	}

	return res, nil
}
