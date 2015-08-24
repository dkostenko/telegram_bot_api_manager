// Package telegram_bot_api provides methods for communicating with Telegram Bot API.
package tgbot

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type (
	SendPhotoRequest struct {
		ChatId           int
		Caption          string
		PhotoUrl         string
		PhotoPath        string
		Photo            string
		ReplyToMessageId int
		ReplyMarkup
	}

	SendPhotoResponse struct {
		Ok          bool    `json:"ok"`
		Result      Message `json:"result"`
		ErrorCode   int     `json:"error_code"`
		Description string  `json:"description"`
	}
)

func (this *SendPhotoRequest) Send(secret string) (*SendPhotoResponse, error) {
	url := getUrl("sendPhoto", secret)

	params, err := this.getParams()
	if err != nil {
		return nil, err
	}

	respJson, err := sendPost(url, params)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(strings.NewReader(respJson))

	var resp *SendPhotoResponse
	if err := dec.Decode(&resp); err == io.EOF {
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *SendPhotoRequest) getParams() ([]*Param, error) {
	replyMarkupJson, _ := json.Marshal(this.ReplyMarkup)

	res := []*Param{
		{Type: "string", Key: "chat_id", Value: strconv.Itoa(this.ChatId)},
		{Type: "string", Key: "reply_markup", Value: string(replyMarkupJson)},
		{Type: "string", Key: "reply_to_message_id", Value: string(this.ReplyToMessageId)},
		{Type: "string", Key: "caption", Value: this.Caption},
	}

	if this.PhotoPath != "" {
		res = append(res, &Param{Type: "filePath", Key: "photo", Value: this.PhotoPath})
	}

	if this.PhotoUrl != "" {
		res = append(res, &Param{Type: "fileUrl", Key: "photo", Value: this.PhotoUrl})
	}

	if this.Photo != "" {
		res = append(res, &Param{Type: "string", Key: "photo", Value: this.Photo})
	}

	return res, nil
}
