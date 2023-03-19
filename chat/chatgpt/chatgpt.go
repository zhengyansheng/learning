package chatgpt

import (
	"context"
	"math/rand"
	"time"
	
	openai "github.com/sashabaranov/go-openai"
)

type chatGPT struct {
	client   *openai.Client
	model    string
	timeout  time.Duration
	proxyUrl string
}

// Option FUNCTIONAL OPTIONS
type Option func(c *chatGPT)

func WithModel(model string) Option {
	return func(c *chatGPT) {
		c.model = model
	}
}

func WithTimeout(timeout int) Option {
	return func(c *chatGPT) {
		c.timeout = time.Second * time.Duration(timeout)
	}
}

func WithProxyUrl(proxyUrl string) Option {
	return func(c *chatGPT) {
		c.proxyUrl = proxyUrl
	}
}

func NewGPT(apiAuthTokens []string, opts ...Option) *chatGPT {
	chat := &chatGPT{
		model:   openai.GPT3Dot5Turbo,                                           // 提供一个默认值
		client:  openai.NewClient(apiAuthTokens[rand.Intn(len(apiAuthTokens))]), // 随机token
		timeout: time.Second * 30,                                               // 默认超时30秒
	}
	for _, opt := range opts {
		opt(chat)
	}
	return chat
}

func (c *chatGPT) Request(content string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	resp, err := c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: c.model,
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
