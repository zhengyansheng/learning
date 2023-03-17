package chatgpt

import (
	"context"
	"math/rand"
	"time"
	
	openai "github.com/sashabaranov/go-openai"
)

type chatGPT struct {
	Client *openai.Client
	Model  string `json:"model"`
}

// Option FUNCTIONAL OPTIONS
type Option func(c *chatGPT)

func WithModel(model string) Option {
	return func(c *chatGPT) {
		c.Model = model
	}
}

func NewGPT(apiAuthTokens []string, opts ...Option) *chatGPT {
	chat := &chatGPT{
		Model:  openai.GPT3Dot5Turbo,                                           // 提供一个默认值
		Client: openai.NewClient(apiAuthTokens[rand.Intn(len(apiAuthTokens))]), // 随机token
	}
	for _, opt := range opts {
		opt(chat)
	}
	return chat
}

func (g *chatGPT) Request(content string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := g.Client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)
	
	if err != nil {
		return "", err
	}
	
	return resp.Choices[0].Message.Content, nil
}
