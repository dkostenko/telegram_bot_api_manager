package tgbot

import (
	"encoding/json"
	"io"
	"strings"
)

type (
	SetWebhookRequest struct {
		Url string
	}

	SetWebhookResponse struct {
		Ok          bool   `json:"ok"`
		Result      bool   `json:"result"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	}
)

func (this *SetWebhookRequest) Send(secret string) (*SetWebhookResponse, error) {
	url := getUrl("setWebhook", secret)

	params, err := this.getParams()
	if err != nil {
		return nil, err
	}

	respJson, err := sendPost(url, params)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(strings.NewReader(respJson))

	var resp *SetWebhookResponse
	if err := dec.Decode(&resp); err == io.EOF {
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *SetWebhookRequest) getParams() ([]*Param, error) {
	res := []*Param{
		{Type: "string", Key: "url", Value: this.Url},
	}

	return res, nil
}
