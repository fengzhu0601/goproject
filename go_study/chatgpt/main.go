package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func main() {
	endPoint := "https://service-8rxb2kbc-1252560870.usw.apigw.tencentcs.com/v1/chat/completions"

	client := resty.New()
	apiKey := "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"

	message := []map[string]string{
		{"role": "user", "content": "这是我第一次使用API"},
	}

	reqBody := make(map[string]interface{})
	reqBody["model"] = "gpt-3.5-turbo"
	reqBody["message"] = message
	reqBody["temperature"] = 1

	reqJSON, _ := json.Marshal(reqBody)

	resp, err := (client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", apiKey)).
		SetBody(string(reqJSON)).
		Post(endPoint))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", resp.String())

}
