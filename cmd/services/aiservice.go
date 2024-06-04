package services

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

type AIService struct {
}

func (aiService AIService) SendSimpleRequest(task string) (string, error) {

	apiKey := os.Getenv("OPENAIKEY")

	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: task,
			},
		},
	})

	answer := resp.Choices[0].Message.Content

	return answer, err
}
