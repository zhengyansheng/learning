package chatgpt

import (
	"context"
	"math/rand"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

type chatGPT struct {
	client  *openai.Client
	model   string
	timeout time.Duration
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

// DefaultSystemChatCompletionMessage chat role: system
func DefaultSystemChatCompletionMessage() openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: "你是一个AI助手，我需要你模拟一名拥有15年经验的DBA工程师来回答我的问题",
	}
}

type ChatResponse struct {
	Reply   string                         `json:"reply"`   // 回复
	Context []openai.ChatCompletionMessage `json:"context"` // 上下文
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

func (c *chatGPT) Request(messages []openai.ChatCompletionMessage) (ChatResponse, error) {
	var newMessages []openai.ChatCompletionMessage

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	if messages[0].Role != "system" {
		newMessages = append(newMessages, DefaultSystemChatCompletionMessage())
	}

	resp, err := c.client.CreateChatCompletion(ctx,
		openai.ChatCompletionRequest{
			Model:    c.model,
			Messages: messages,
		},
	)

	if err != nil {
		return ChatResponse{}, err
	}

	return ChatResponse{
		Reply:   resp.Choices[0].Message.Content,
		Context: append(newMessages, resp.Choices[0].Message),
	}, nil
}

func (c *chatGPT) SimpleRequest(content string) (ChatResponse, error) {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		},
	}
	return c.Request(messages)
}
