// Package telegram_bot_api provides methods for communicating with Telegram Bot API.
package tgbot

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type (
	SendAudioRequest struct {
		ChatId           int
		Title            string
		Performer        string
		AudioUrl         string
		AudioPath        string
		Audio            string
		ReplyToMessageId int
		Duration         int
		ReplyMarkup
	}

	SendAudioResponse struct {
		Ok          bool    `json:"ok"`
		Result      Message `json:"result"`
		ErrorCode   int     `json:"error_code"`
		Description string  `json:"description"`
	}
)

func (this *SendAudioRequest) Send(secret string) (*SendAudioResponse, error) {
	url := getUrl("sendAudio", secret)

	params, err := this.getParams()
	if err != nil {
		return nil, err
	}

	respJson, err := sendPost(url, params)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(strings.NewReader(respJson))

	var resp *SendAudioResponse
	if err := dec.Decode(&resp); err == io.EOF {
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *SendAudioRequest) getParams() ([]*Param, error) {
	replyMarkupJson, _ := json.Marshal(this.ReplyMarkup)

	res := []*Param{
		{Type: "string", Key: "chat_id", Value: strconv.Itoa(this.ChatId)},
		{Type: "string", Key: "reply_markup", Value: string(replyMarkupJson)},
		{Type: "string", Key: "reply_to_message_id", Value: string(this.ReplyToMessageId)},
		{Type: "string", Key: "duration", Value: strconv.Itoa(this.Duration)},
		{Type: "string", Key: "title", Value: this.Title},
		{Type: "string", Key: "performer", Value: this.Performer},
	}

	if this.AudioPath != "" {
		res = append(res, &Param{Type: "filePath", Key: "audio", Value: this.AudioPath})
	}

	if this.AudioUrl != "" {
		res = append(res, &Param{Type: "fileUrl", Key: "audio", Value: this.AudioUrl})
	}

	if this.Audio != "" {
		res = append(res, &Param{Type: "string", Key: "audio", Value: this.Audio})
	}

	return res, nil
}
