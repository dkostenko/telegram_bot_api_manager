// Package telegram_bot_api provides methods for communicating with Telegram Bot API.
package tgbot

import (
    // "log"
    "strconv"
    "io"
    "strings"
    "encoding/json"
)

type SendMessageRequest struct {
    ChatId int
    Text string
    ReplyToMessageId int
    ReplyMarkup
}

type ReplyMarkup struct {
    Keyboard [][]string `json:"keyboard"`
    ResizeKeyboard bool `json:"resize_keyboard"`
    OneTimeKeyboard bool `json:"one_time_keyboard"`
    Selective bool `json:"selective"`
    HideKeyboard bool `json:"hide_keyboard"`
    ForceReply bool `json:"force_reply"`
}

func (this *SendMessageRequest) Send() (*SendMessageResponse, error){
    url := getUrl("sendMessage")
    
    params, err := this.getParams()
    if err != nil { return nil, err }
    
    respJson, err := sendPost(url, params)
    if err != nil { return nil, err}
    
    dec := json.NewDecoder(strings.NewReader(respJson))
    
    var resp *SendMessageResponse
	if err := dec.Decode(&resp); err == io.EOF {
		return resp, nil
	} else if err != nil {
        return nil, err
	}
    
    return resp, nil
}

func (this *SendMessageRequest) getParams() ([]*Param, error){
    replyMarkupJson, _ := json.Marshal(this.ReplyMarkup)
    
    res := []*Param{
        &Param{Key: "chat_id", Value: strconv.Itoa(this.ChatId)}, 
        &Param{Key: "text", Value: this.Text},
        &Param{Key: "reply_markup", Value: string(replyMarkupJson)},
        &Param{Key: "reply_to_message_id", Value: string(this.ReplyToMessageId)},
    }
    
    return res, nil
}