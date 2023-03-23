package chatgpt

import (
	"context"
	"math/rand"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

var (
	sqlStatementOptimization = "优化建议：1. 当只要一行数据时使用 LIMIT 1, 2. 为搜索字段建索引, 3. 在Join表的时候使用相当类型的例，并将其索引, 4. 千万不要 ORDER BY RAND(), 5. 避免 SELECT *, 6. 永远为每张表设置一个ID, 7. 使用 ENUM 而不是 VARCHAR, 8. 对查询进行优化，要尽量避免全表扫描, 9. 应尽量避免在 where 子句中对字段进行 null 值判断, 10. 应尽量避免在 where 子句中使用 != 或 <> 操作符, 11. 应尽量避免在 where 子句中使用 or 来连接条件，否则系统将进行全表扫描, 12. 模糊查询前边不要使用%，否则系统将可能无法正确使用索引, 13. 应尽量避免在where子句中对字段进行函数操作，否则系统将可能无法正确使用索引, 14. 不要在 where 子句中的=左边进行函数、算术运算或其他表达式运算，否则系统将可能无法正确使用索引, 15. 不带条件的查询是全表扫描，不允许发起, 16. 尽量避免大事务操作，提高系统并发能力"
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
		Content: "你是一个乐于助人的AI助手，我需要你模拟一位拥有15年经验的MySQL工程师来回答我的问题",
	}
}

// DefaultSQLChatCompletionMessage MySQL optimization
func DefaultSQLChatCompletionMessage() []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "你是一个乐于助人的AI助手，我需要你模拟一位拥有15年经验的MySQL工程师来回答我的问题",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "你是谁?",
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "我是一名AI助手，拥有15年经验的MySQL工程师，很高兴为你服务，请你说出你的SQL语句，我将要优化，提建议和优化后的SQL语句",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: sqlStatementOptimization,
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "我已经学会了SQL优化标准，接下来SQL语句我将按照以上标准执行优化, 建议和提供优化后的SQL",
		},
	}

}

type Response struct {
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

// Request /chat/completions
func (c *chatGPT) Request(messages []openai.ChatCompletionMessage) (Response, error) {
	var newMessages []openai.ChatCompletionMessage

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	if messages[0].Role != "system" {
		newMessages = append(newMessages, DefaultSystemChatCompletionMessage())
	}

	resp, err := c.client.CreateChatCompletion(ctx,
		openai.ChatCompletionRequest{
			Model:            c.model,
			Messages:         messages,
			PresencePenalty:  1, // 默认值是0, 惩罚机制，默认是-2.0 到2.0 之间, 数值越小提交重复令牌数越多，从而能更清楚文本的意思
			FrequencyPenalty: 1, // 默认值是0, 频率，默认是-2.0 到2.0 之间
			Temperature:      1, // 默认值是1, 温度，默认是0-2，调整回复的精确度使用
			N:                1, // 默认条数
		},
	)

	if err != nil {
		return Response{}, err
	}

	return Response{
		Reply:   resp.Choices[0].Message.Content,
		Context: append(newMessages, resp.Choices[0].Message),
	}, nil
}

func (c *chatGPT) SimpleRequest(content string) (Response, error) {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		},
	}
	return c.Request(messages)
}
