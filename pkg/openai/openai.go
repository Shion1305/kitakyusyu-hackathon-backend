package openai

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"kitakyusyu-hackathon/pkg/config"
)

type OpenAI struct {
	endpoint string
	apiKey   string
	client   *openai.Client
}

func NewOpenAI() *OpenAI {
	conf := config.Get()
	return &OpenAI{
		endpoint: conf.OpenAI.EndPoint,
		apiKey:   conf.OpenAI.APIKey,
		client:   openai.NewClient(conf.OpenAI.APIKey),
	}
}

func (o OpenAI) GetStreamResponse(ctx context.Context, input []string) (*openai.ChatCompletionStream, error) {
	messages := make([]openai.ChatCompletionMessage, 0, len(input)+1)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: "You are an assistant helping a user with a task.",
	})
	for i := range input {
		var role string
		if i%2 == 0 {
			role = openai.ChatMessageRoleUser
		} else {
			role = openai.ChatMessageRoleAssistant
		}
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    role,
			Content: input[i],
		})
	}
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 20,
		Messages:  messages,
		Stream:    true,
	}

	stream, err := o.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return nil, fmt.Errorf("failed to create chat completion stream: %w", err)
	}

	return stream, nil
}
