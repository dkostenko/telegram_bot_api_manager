package tgbot

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type (
	SendVideoRequest struct {
		ChatId           int
		Caption          string
		VideoUrl         string
		VideoPath        string
		Video            string
		ReplyToMessageId int
		Duration         int
		ReplyMarkup
	}

	SendVideoResponse struct {
		Ok          bool    `json:"ok"`
		Result      Message `json:"result"`
		ErrorCode   int     `json:"error_code"`
		Description string  `json:"description"`
	}
)

func (this *SendVideoRequest) Send(secret string) (*SendVideoResponse, error) {
	url := getUrl("sendVideo", secret)

	params, err := this.getParams()
	if err != nil {
		return nil, err
	}

	respJson, err := sendPost(url, params)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(strings.NewReader(respJson))

	var resp *SendVideoResponse
	if err := dec.Decode(&resp); err == io.EOF {
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *SendVideoRequest) getParams() ([]*Param, error) {
	replyMarkupJson, _ := json.Marshal(this.ReplyMarkup)

	res := []*Param{
		{Type: "string", Key: "chat_id", Value: strconv.Itoa(this.ChatId)},
		{Type: "string", Key: "reply_markup", Value: string(replyMarkupJson)},
		{Type: "string", Key: "reply_to_message_id", Value: string(this.ReplyToMessageId)},
		{Type: "string", Key: "duration", Value: strconv.Itoa(this.Duration)},
		{Type: "string", Key: "caption", Value: this.Caption},
	}

	if this.VideoPath != "" {
		res = append(res, &Param{Type: "filePath", Key: "video", Value: this.VideoPath})
	}

	if this.VideoUrl != "" {
		res = append(res, &Param{Type: "fileUrl", Key: "video", Value: this.VideoUrl})
	}

	if this.Video != "" {
		res = append(res, &Param{Type: "string", Key: "video", Value: this.Video})
	}

	return res, nil
}
