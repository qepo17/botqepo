package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type CompletionsResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type CompletionsRequest struct {
	Model     string  `json:"model"`
	Prompt    string  `json:"prompt"`
	TopP      float32 `json:"top_p"`
	MaxTokens int     `json:"max_tokens"`
}

type CfgOpenai struct {
	BaseUrl string
	Token   string
	Model   string
}

func Completions(cfg CfgOpenai, keyword string) CompletionsResponse {
	keywordArray := strings.Split(keyword, " ")
	var filteredKeyword string
	for _, value := range keywordArray {
		if !strings.Contains(value, "@") {
			filteredKeyword = fmt.Sprintf("%s %s", filteredKeyword, value)
		}

		filteredKeyword = strings.TrimSpace(filteredKeyword)
	}

	logicToMakeAnswerBetter := "Q: What is human life expectancy in the United States?\nA: Human life expectancy in the United States is 78 years.\n\nQ: Where were the 1992 Olympics held?\nA: The 1992 Olympics were held in Barcelona, Spain.\n\n"
	request := CompletionsRequest{
		Model:     cfg.Model,
		Prompt:    fmt.Sprintf("%s%s", logicToMakeAnswerBetter, "Q: "+filteredKeyword+"\n"),
		TopP:      1.0,
		MaxTokens: 30,
	}

	byteReq, err := json.Marshal(request)
	if err != nil {
		log.Fatalln(err)
	}

	bodyReq := bytes.NewReader([]byte(byteReq))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/completions", cfg.BaseUrl), bodyReq)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+cfg.Token)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.Body.Close()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	var result CompletionsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}
