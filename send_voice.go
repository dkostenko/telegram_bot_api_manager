package tgbot

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type (
	SendVoiceRequest struct {
		ChatId           int
		AudioUrl         string
		AudioPath        string
		Audio            string
		ReplyToMessageId int
		Duration         int
		ReplyMarkup
	}

	SendVoiceResponse struct {
		Ok          bool    `json:"ok"`
		Result      Message `json:"result"`
		ErrorCode   int     `json:"error_code"`
		Description string  `json:"description"`
	}
)

func (this *SendVoiceRequest) Send(secret string) (*SendVoiceResponse, error) {
	url := getUrl("sendVoice", secret)

	params, err := this.getParams()
	if err != nil {
		return nil, err
	}

	respJson, err := sendPost(url, params)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(strings.NewReader(respJson))

	var resp *SendVoiceResponse
	if err := dec.Decode(&resp); err == io.EOF {
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *SendVoiceRequest) getParams() ([]*Param, error) {
	replyMarkupJson, _ := json.Marshal(this.ReplyMarkup)

	res := []*Param{
		{Type: "string", Key: "chat_id", Value: strconv.Itoa(this.ChatId)},
		{Type: "string", Key: "reply_markup", Value: string(replyMarkupJson)},
		{Type: "string", Key: "reply_to_message_id", Value: string(this.ReplyToMessageId)},
		{Type: "string", Key: "duration", Value: strconv.Itoa(this.Duration)},
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
