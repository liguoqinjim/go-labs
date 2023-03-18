package main

import (
	"context"
	"errors"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("需要输入APIKEY")
	}
	apiKey := os.Args[1]

	c := openai.NewClient(apiKey)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo0301,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "帮忙写一篇30字左右的周报",
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("CompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			return
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			return
		}

		fmt.Printf("Stream response: %v\n", response)
	}
}
