package chat

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

type chatGPT struct {
	client    *openai.Client // open AI client
	model     string         // open AI model(GPT-3.4 / GPT-4/ ...)
	timeout   time.Duration  // request timeout
	MaxTokens int
	Stream    bool
}

type ImagePrompt struct {
	Prompt string `json:"prompt"` // 提示语
	Size   string `json:"size"`   // 图片大小
	Nums   int    `json:"nums"`   // 图片数量
}

type Response struct {
	Reply   string                         `json:"reply"`   // 回复
	Context []openai.ChatCompletionMessage `json:"context"` // 上下文
}

// StreamMessage chat reply with stream
type streamMessage struct {
	Ch  chan string // chat text stream
	Err error       // error message
}

func NewGPT(apiKey, baseURL string, opts ...Option) *chatGPT {
	cfg := openai.DefaultAzureConfig(apiKey, baseURL)
	chat := &chatGPT{
		model:   openai.GPT3Dot5Turbo,            // 提供一个默认值
		client:  openai.NewClientWithConfig(cfg), // 随机token
		timeout: time.Second * 60,                // 默认超时60秒
	}
	for _, opt := range opts {
		opt(chat)
	}
	return chat
}

// CreateChatCompletionWithContent /chat/completions
func (c *chatGPT) CreateChatCompletionWithContent(messages []openai.ChatCompletionMessage) (Response, error) {
	var newMessages []openai.ChatCompletionMessage
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.client.CreateChatCompletion(ctx,
		openai.ChatCompletionRequest{
			Model:     c.model,
			Messages:  messages,
			MaxTokens: c.MaxTokens,
			//PresencePenalty:  1, // 默认值是0, 惩罚机制，默认是-2.0 到2.0 之间, 数值越小提交重复令牌数越多，从而能更清楚文本的意思
			//FrequencyPenalty: 1, // 默认值是0, 频率，默认是-2.0 到2.0 之间
			//Temperature:      1, // 默认值是1, 温度，默认是0-2，调整回复的精确度使用
			//N:                1, // 默认条数
		},
	)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return Response{}, fmt.Errorf("timeout: %v", "会话请求超时, 请再次尝试")
		}
		return Response{}, err
	}

	return Response{
		Reply:   resp.Choices[0].Message.Content,
		Context: append(newMessages, resp.Choices[0].Message),
	}, nil

}

// CreateChatCompletionStream /chat/completions stream; 支持 gpt-3.5 ...
func (c *chatGPT) CreateChatCompletionStream(messages []openai.ChatCompletionMessage) (*streamMessage, error) {
	req := openai.ChatCompletionRequest{
		Model:     c.model,
		Messages:  messages,
		MaxTokens: c.MaxTokens,
		Stream:    c.Stream,
	}

	streamMsg := &streamMessage{Ch: make(chan string), Err: nil}
	stream, err := c.client.CreateChatCompletionStream(context.TODO(), req)
	if err != nil {
		return streamMsg, err
	}

	go func() {
		defer stream.Close()
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				streamMsg.Err = fmt.Errorf("stream finished")
				break
			}

			if err != nil {
				streamMsg.Err = fmt.Errorf("stream error: %v", err)
				break
			}
			for _, choice := range response.Choices {
				streamMsg.Ch <- choice.Delta.Content
			}
		}
		streamMsg.Ch <- "[DONE]"
		close(streamMsg.Ch)
		return
	}()

	return streamMsg, nil
}

// CreateChatStream /v1/completions; 不支持 gpt-3.5
func (c *chatGPT) CreateChatStream(prompt string) (*streamMessage, error) {
	req := openai.CompletionRequest{
		Model:       c.model,
		Prompt:      prompt,
		MaxTokens:   c.MaxTokens,
		Stream:      c.Stream,
		Temperature: 0,
	}

	streamMsg := &streamMessage{Ch: make(chan string), Err: nil}
	stream, err := c.client.CreateCompletionStream(context.TODO(), req)
	if err != nil {
		return streamMsg, err
	}

	go func() {
		// TODO: here, so not data
		defer stream.Close()
		defer close(streamMsg.Ch)
		defer func() {
			streamMsg.Ch <- "[DONE]"
		}()

		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				streamMsg.Err = fmt.Errorf("stream finished")
				return
			}

			if err != nil {
				streamMsg.Err = fmt.Errorf("stream error: %v", err)
				return
			}
			for _, choice := range response.Choices {
				streamMsg.Ch <- choice.Text
			}
		}
	}()

	return streamMsg, nil
}

// CreateImage /images/generations
func (c *chatGPT) CreateImage(img ImagePrompt) (images []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	req := openai.ImageRequest{
		Prompt: img.Prompt,
		N:      img.Nums,
		Size:   img.Size,
	}

	resp, err := c.client.CreateImage(ctx, req)
	if err != nil {
		return
	}

	for _, img := range resp.Data {
		images = append(images, img.URL)
	}
	return
}

// SimpleChatCompletion simple chat completion
func (c *chatGPT) SimpleChatCompletion(content string) (Response, error) {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		},
	}
	return c.CreateChatCompletionWithContent(messages)
}
