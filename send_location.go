package tgbot

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type (
	SendLocationRequest struct {
		ChatId           int
		Latitude         float64
		Longitude        float64
		ReplyToMessageId int
		ReplyMarkup
	}

	SendLocationResponse struct {
		Ok          bool    `json:"ok"`
		Result      Message `json:"result"`
		ErrorCode   int     `json:"error_code"`
		Description string  `json:"description"`
	}
)

func (this *SendLocationRequest) Send(secret string) (*SendLocationResponse, error) {
	url := getUrl("sendLocation", secret)

	params, err := this.getParams()
	if err != nil {
		return nil, err
	}

	respJson, err := sendPost(url, params)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(strings.NewReader(respJson))

	var resp *SendLocationResponse
	if err := dec.Decode(&resp); err == io.EOF {
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *SendLocationRequest) getParams() ([]*Param, error) {
	replyMarkupJson, _ := json.Marshal(this.ReplyMarkup)

	res := []*Param{
		{Type: "string", Key: "chat_id", Value: strconv.Itoa(this.ChatId)},
		{Type: "string", Key: "latitude", Value: strconv.FormatFloat(this.Latitude, 'f', 6, 64)},
		{Type: "string", Key: "longitude", Value: strconv.FormatFloat(this.Longitude, 'f', 6, 64)},
		{Type: "string", Key: "reply_markup", Value: string(replyMarkupJson)},
		{Type: "string", Key: "reply_to_message_id", Value: string(this.ReplyToMessageId)},
	}

	return res, nil
}
