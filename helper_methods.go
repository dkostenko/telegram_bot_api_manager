// Package tgbot provides methods for communicating with Telegram Bot API.
package tgbot

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func sendGet(url string, params []*Param) (string, error) {
	url += "?"

	for i, param := range params {
		if i > 0 {
			url += "&"
		}
		url += param.Key + "=" + param.Value
	}

	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	resBody, _ := ioutil.ReadAll(resp.Body)
	json := string(resBody)

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	log.Println("response Body:", json)

	return json, nil
}

func sendPost(url string, params []*Param) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for _, param := range params {
		switch param.Type {
		default:
			return "", errors.New("No type in param")
		case "string":
			writer.WriteField(param.Key, param.Value)
		case "fileUrl":
			res, err := http.Get(param.Value)

			if err != nil {
				return "", err
			}
			defer res.Body.Close()

			file := res.Body
			part, err := writer.CreateFormFile(param.Key, filepath.Base(param.Value))
			if err != nil {
				return "", err
			}

			_, err = io.Copy(part, file)
			if err != nil {
				return "", err
			}
		case "filePath":
			file, err := os.Open(param.Value)
			if err != nil {
				return "", err
			}

			part, err := writer.CreateFormFile(param.Key, filepath.Base(param.Value))
			if err != nil {
				return "", err
			}

			_, err = io.Copy(part, file)
			if err != nil {
				return "", err
			}
		}
	}

	err := writer.Close()
	if err != nil {
		return "", err
	}

	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	resBody, _ := ioutil.ReadAll(resp.Body)
	json := string(resBody)

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	log.Println("response Body:", json)

	return json, nil
}

func getUrl(method string, secret string) (url string) {
	url = "https://api.telegram.org/" + secret + "/" + method
	return
}
