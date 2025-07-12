package service

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"go-ai-summarizer/config"
	"io/ioutil"
	"net/http"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type DeepSeekRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type DeepSeekResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

func CallDeepSeekAPI(input string) (string, error) {
	reqBody := DeepSeekRequest{
		Model: "deepseek-chat",
		Messages: []ChatMessage{
			{Role: "user", Content: "请用简洁中文总结以下内容：\n" + input},
		},
	}

	jsonData, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "https://api.deepseek.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+config.GetDeepSeekKey())
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", errors.New(fmt.Sprintf("DeepSeek API error: %s", body))
	}

	var result DeepSeekResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	if len(result.Choices) == 0 {
		return "", errors.New("no summary returned")
	}

	return result.Choices[0].Message.Content, nil
}

// mock
//func CallDeepSeekAPI(input string) (string, error) {
//	// Mock 逻辑：返回固定字符串
//	return "这是模拟的摘要结果：内容大致表达了..." + input[:min(len(input), 500)] + "...", nil
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
