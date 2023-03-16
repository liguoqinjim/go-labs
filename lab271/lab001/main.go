package main

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("需要输入APIKEY")
	}
	apiKey := os.Args[1]

	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo0301,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return
	}

	log.Println(resp.Choices[0].Message.Content)
}
