package chat_gpt

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"test/helpers"
)

func ChatGPT() {
	// Cấu hình thông tin yêu cầu
	apiURL := "https://api.openai.com/v1/chat/completions"
	apiKey := "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"

	a := []map[string]interface{}{}
	a = append(a, map[string]interface{}{
		"role":    "user",
		"content": "merkle tree là gì",
	})
	// Dữ liệu yêu cầu
	requestData := map[string]interface{}{
		"model":       "gpt-3.5-turbo",
		"messages":    a,
		"temperature": 0.7,
	}

	// Chuyển đổi dữ liệu yêu cầu thành JSON
	jsonData, _ := json.Marshal(requestData)

	b, err := helpers.SendRequest("POST", apiURL, fiber.Map{
		"Authorization": "Bearer " + apiKey,
		"Content-Type":  "application/json",
	}, jsonData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}
