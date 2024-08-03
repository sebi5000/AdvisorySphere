package services

import (
	"context"
	"errors"
	"sphere/cmd/model/security"

	"github.com/sashabaranov/go-openai"
)

type AIService struct {
}

func (aiService AIService) SendPromptedRequest(preprompt string, task string) (string, error) {

	apiKey := security.GetInstance().OPENAIKEY

	if len(apiKey) == 0 {
		return "", errors.New("please set api-key for using ai features")
	}

	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: preprompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: task,
			},
		},
	})

	answer := resp.Choices[0].Message.Content

	return answer, err
}

func (aiService AIService) SendSimpleRequest(task string) (string, error) {
	return aiService.SendPromptedRequest("", task)
}
