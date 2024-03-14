package helpers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"math/rand"
	"net"
	"time"
)

func SendRequest(method, url string, header fiber.Map, data []byte) (response []byte, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	for key, value := range header {
		req.Header.Set(key, value.(string))
	}
	req.Header.SetMethod(method)
	req.Header.SetContentType("application/json")

	if method == "POST" || method == "PUT" {
		req.SetBody(data)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	//start := time.Now()

	// new fix
	client := &fasthttp.Client{
		ReadTimeout:         time.Duration(60) * time.Second,
		MaxIdleConnDuration: time.Duration(600) * time.Second,
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, time.Duration(60)*time.Second)
		},
		TLSConfig:          &tls.Config{InsecureSkipVerify: true},
		MaxConnsPerHost:    4,
		MaxConnWaitTimeout: time.Minute,
	}
	errors := client.DoTimeout(req, resp, time.Duration(60)*time.Second)
	if errors != nil {
		fmt.Println("Url:", url, string(data))
		fmt.Printf("Client get failed: %s\n", errors)
		return nil, errors
	}

	var body []byte
	body = resp.Body()

	return body, errors
}

func SendRequestTmp(method, url string, header fiber.Map, data []byte) (response []byte, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	for key, value := range header {
		req.Header.Set(key, value.(string))
	}
	req.Header.SetMethod(method)
	req.Header.SetContentType("application/json")

	if method == "POST" {
		req.SetBody(data)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	client := &fasthttp.Client{
		ReadTimeout:         time.Duration(60) * time.Second,
		MaxIdleConnDuration: time.Duration(600) * time.Second,
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, time.Duration(60)*time.Second)
		},
		TLSConfig:          &tls.Config{InsecureSkipVerify: true},
		MaxConnsPerHost:    4,
		MaxConnWaitTimeout: time.Minute,
	}
	errors := client.DoTimeout(req, resp, time.Duration(60)*time.Second)

	if errors != nil {
		fmt.Println("Url:", url, string(data))
		fmt.Printf("Client get failed: %s\n", errors)
		return nil, errors
	}

	var body []byte
	body = resp.Body()
	bodyRes := map[string]interface{}{}
	_ = json.Unmarshal(body, &bodyRes)
	if bodyRes["status"] != nil && bodyRes["data"] != nil {
		if bodyRes["status"].(bool) == true {
			bodyResDataJson, _ := json.Marshal(bodyRes["data"].(map[string]interface{}))
			return bodyResDataJson, errors
		} else {
			return nil, errors
		}
	} else {
		return nil, errors
	}
}

func Random(min int, max int) int {
	return rand.Intn(max-min) + min
}
