// Package telegram_bot_api provides methods for communicating with Telegram Bot API.
package telegram_bot_api

import (
    "sync"
    "log"
    "net/http"
    // "net/url"
    "io/ioutil"
    "strconv"
    "mime/multipart"
    "bytes"
    "path/filepath"
    "io"
    // "reflect"
)

type (
    Param struct {
        Key string
        Value string
        FileUrl string
    }

    ApiManager struct {
        baseApi string
        secretToken string
        DebugMode bool
    }
)

var (
    once sync.Once
    instance *ApiManager
)

// NewApiManager returns a singleton API manager.
func NewApiManager(baseApi string, secretToken string, debugMode bool) *ApiManager {
    once.Do(func() {
        instance = &ApiManager{
            baseApi: baseApi,
            secretToken: secretToken,
            DebugMode: debugMode,
        }
    })
    return instance
}

func (this *ApiManager) SendMessage(chatId int, text string, replyToMessageId int) {
    url := this.getUrl("sendMessage")

    params := []*Param{
        &Param{Key: "text", Value: text},
        &Param{Key: "chat_id", Value: strconv.Itoa(chatId)},
        &Param{Key: "reply_markup", Value: "{\"resize_keyboard\":true,\"one_time_keyboard\":true,\"keyboard\":[[\"first\",\"second\"],[\"4444\",\"5555\"]]"},
        // &Param{Key: "reply_markup", Value: "{\"force_reply\":true}"},
    }
    
    if replyToMessageId > 0 {
        params = append(params, &Param{Key: "reply_to_message_id", Value: strconv.Itoa(replyToMessageId)})
    }

    _, err := this.sendPost(url, params)
    if err != nil {
        log.Println("ERROR", err)
    }
}

func (this *ApiManager) GetMe() {
    url := this.getUrl("getMe")

    this.sendGet(url, nil)
}

func (this *ApiManager) SendPhoto(chatId int, fileUrl string, caption string, replyToMessageId int) {
    url := this.getUrl("sendPhoto")

    params := []*Param{
        &Param{Key: "photo", FileUrl: fileUrl},
        &Param{Key: "chat_id", Value: strconv.Itoa(chatId)},
    }
    
    if caption != "" {
        params = append(params, &Param{Key: "caption", Value: caption},)
    }
    
    if replyToMessageId > 0 {
        params = append(params, &Param{Key: "reply_to_message_id", Value: strconv.Itoa(replyToMessageId)})
    }

    this.sendPost(url, params)
}

func (this *ApiManager) SetWebhook(webhookUrl string) {
    url := this.getUrl("setWebhook")
    
    params := []*Param{
        &Param{Key: "url", Value: webhookUrl},
    }
    
    this.sendGet(url, params)
}

func (this *ApiManager) sendPost(url string, params []*Param) ([]int, error) {
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    
    for _, param := range params {
        if param.FileUrl != "" {
            res, err := http.Get(param.FileUrl)
            
            if err != nil { return nil, err }
            defer res.Body.Close()
            
            file := res.Body
            part, err := writer.CreateFormFile(param.Key, filepath.Base(param.FileUrl))
            if err != nil { return nil, err }
            
            _, err = io.Copy(part, file)
            if err != nil { return nil, err }
        } else {
            writer.WriteField(param.Key, param.Value)
        }
    }

    err := writer.Close()
    if err != nil { return nil, err }

    req, _ := http.NewRequest("POST", url, body)
    req.Header.Set("Content-Type", writer.FormDataContentType())
    
    client := &http.Client{}
    resp, err := client.Do(req)
    
    if err != nil { return nil, err }
    defer resp.Body.Close()

    resBody, _ := ioutil.ReadAll(resp.Body)

    if this.DebugMode {
        log.Println("response Status:", resp.Status)
        log.Println("response Headers:", resp.Header)
        log.Println("response Body:", string(resBody))
    }
    
    return nil, nil
}

func (this *ApiManager) sendGet(url string, params []*Param) {
    url += "?"
    
    for i, param := range params {
        if i > 0 { url += "&" }
        url += param.Key + "=" + param.Value
    }
    
    resp, err := http.Get(url)
    
    if err != nil {
    	log.Println(err)
    }
    
    defer resp.Body.Close()
    
    body, _ := ioutil.ReadAll(resp.Body)
    
    if this.DebugMode {
        log.Println("response Status:", resp.Status)
        log.Println("response Headers:", resp.Header)
        log.Println("response Body:", string(body))
    }
}

func (this *ApiManager) getUrl(method string) (url string) {
    url = this.baseApi + "/" + this.secretToken + "/" + method
    return
}
