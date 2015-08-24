// Package telegram_bot_api provides methods for communicating with Telegram Bot API.
package tgbot

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type (
	SendDocumentRequest struct {
		ChatId           int
		Caption          string
		DocumentUrl      string
		DocumentPath     string
		Document         string
		ReplyToMessageId int
		ReplyMarkup
	}

	SendDocumentResponse struct {
		Ok          bool    `json:"ok"`
		Result      Message `json:"result"`
		ErrorCode   int     `json:"error_code"`
		Description string  `json:"description"`
	}
)

func (this *SendDocumentRequest) Send(secret string) (*SendDocumentResponse, error) {
	url := getUrl("sendDocument", secret)

	params, err := this.getParams()
	if err != nil {
		return nil, err
	}

	respJson, err := sendPost(url, params)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(strings.NewReader(respJson))

	var resp *SendDocumentResponse
	if err := dec.Decode(&resp); err == io.EOF {
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *SendDocumentRequest) getParams() ([]*Param, error) {
	replyMarkupJson, _ := json.Marshal(this.ReplyMarkup)

	res := []*Param{
		{Type: "string", Key: "chat_id", Value: strconv.Itoa(this.ChatId)},
		{Type: "string", Key: "reply_markup", Value: string(replyMarkupJson)},
		{Type: "string", Key: "reply_to_message_id", Value: string(this.ReplyToMessageId)},
	}

	if this.DocumentPath != "" {
		res = append(res, &Param{Type: "filePath", Key: "document", Value: this.DocumentPath})
	}

	if this.DocumentUrl != "" {
		res = append(res, &Param{Type: "fileUrl", Key: "document", Value: this.DocumentUrl})
	}

	if this.Document != "" {
		res = append(res, &Param{Type: "string", Key: "document", Value: this.Document})
	}

	return res, nil
}
