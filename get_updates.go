package tgbot

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type (
	GetUpdatesRequest struct {
		Offset  int
		Limit   int
		Timeout int
	}

	GetUpdatesResponse struct {
		Ok          bool      `json:"ok"`
		Result      []*Update `json:"result"`
		ErrorCode   int       `json:"error_code"`
		Description string    `json:"description"`
	}
)

func (this *GetUpdatesRequest) Send(secret string) (*GetUpdatesResponse, error) {
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

	var resp *GetUpdatesResponse
	if err := dec.Decode(&resp); err == io.EOF {
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *GetUpdatesResponse) getParams() ([]*Param, error) {
	res := []*Param{
		{Type: "string", Key: "offset", Value: strconv.Itoa(this.Offset)},
		{Type: "string", Key: "limit", Value: strconv.Itoa(this.Limit)},
		{Type: "string", Key: "timeout", Value: strconv.Itoa(this.Timeout)},
	}

	return res, nil
}
