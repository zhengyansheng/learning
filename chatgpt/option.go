package chatgpt

import "time"

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

func WithMaxTokens(maxTokens int) Option {
	return func(c *chatGPT) {
		c.MaxTokens = maxTokens
	}
}

func WithStream(stream bool) Option {
	return func(c *chatGPT) {
		c.Stream = stream
	}
}
