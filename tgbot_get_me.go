// Package telegram_bot_api provides methods for communicating with Telegram Bot API.
package tgbot

import (
	"encoding/json"
	"io"
	"strings"
)

type (
	GetMeRequest struct {
	}

	GetMeResponse struct {
		Ok          bool   `json:"ok"`
		Result      User   `json:"result"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	}
)

func (this *GetMeRequest) Send(secret string) (*GetMeResponse, error) {
	url := getUrl("getMe", secret)

	params, err := this.getParams()
	if err != nil {
		return nil, err
	}

	respJson, err := sendGet(url, params)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(strings.NewReader(respJson))

	var resp *GetMeResponse
	if err := dec.Decode(&resp); err == io.EOF {
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *GetMeRequest) getParams() ([]*Param, error) {
	return []*Param{}, nil
}
