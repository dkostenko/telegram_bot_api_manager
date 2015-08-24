package tgbot

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type (
	SendStickerRequest struct {
		ChatId           int
		Caption          string
		StickerUrl       string
		StickerPath      string
		Sticker          string
		ReplyToMessageId int
		ReplyMarkup
	}

	SendStickerResponse struct {
		Ok          bool    `json:"ok"`
		Result      Message `json:"result"`
		ErrorCode   int     `json:"error_code"`
		Description string  `json:"description"`
	}
)

func (this *SendStickerRequest) Send(secret string) (*SendStickerResponse, error) {
	url := getUrl("sendSticker", secret)

	params, err := this.getParams()
	if err != nil {
		return nil, err
	}

	respJson, err := sendPost(url, params)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(strings.NewReader(respJson))

	var resp *SendStickerResponse
	if err := dec.Decode(&resp); err == io.EOF {
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *SendStickerRequest) getParams() ([]*Param, error) {
	replyMarkupJson, _ := json.Marshal(this.ReplyMarkup)

	res := []*Param{
		{Type: "string", Key: "chat_id", Value: strconv.Itoa(this.ChatId)},
		{Type: "string", Key: "reply_markup", Value: string(replyMarkupJson)},
		{Type: "string", Key: "reply_to_message_id", Value: string(this.ReplyToMessageId)},
	}

	if this.StickerPath != "" {
		res = append(res, &Param{Type: "filePath", Key: "sticker", Value: this.StickerPath})
	}

	if this.StickerUrl != "" {
		res = append(res, &Param{Type: "fileUrl", Key: "sticker", Value: this.StickerUrl})
	}

	if this.Sticker != "" {
		res = append(res, &Param{Type: "string", Key: "sticker", Value: this.Sticker})
	}

	return res, nil
}
