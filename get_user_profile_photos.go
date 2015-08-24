package tgbot

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type (
	GetUserProfilePhotosRequest struct {
		UserId int
		Offset int
		Limit  int
	}

	GetUserProfilePhotosResponse struct {
		Ok          bool   `json:"ok"`
		Result      bool   `json:"result"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	}
)

func (this *GetUserProfilePhotosRequest) Send(secret string) (*GetUserProfilePhotosResponse, error) {
	url := getUrl("getUserProfilePhotos", secret)

	params, err := this.getParams()
	if err != nil {
		return nil, err
	}

	respJson, err := sendPost(url, params)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(strings.NewReader(respJson))

	var resp *GetUserProfilePhotosResponse
	if err := dec.Decode(&resp); err == io.EOF {
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *GetUserProfilePhotosRequest) getParams() ([]*Param, error) {
	res := []*Param{
		{Type: "string", Key: "user_id", Value: strconv.Itoa(this.UserId)},
		{Type: "string", Key: "offset", Value: strconv.Itoa(this.Offset)},
		{Type: "string", Key: "limit", Value: strconv.Itoa(this.Limit)},
	}

	return res, nil
}
