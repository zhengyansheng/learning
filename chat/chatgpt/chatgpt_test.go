package chatgpt

import (
	"testing"
	
	"github.com/sashabaranov/go-openai"
)

func TestChatGPT(t *testing.T) {
	apiAuthTokens := []string{"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"}
	g := NewGPT(apiAuthTokens, WithModel(openai.GPT3Dot5Turbo0301))
	response, err := g.Request("hello world")
	if err != nil {
		panic(err)
	}
	t.Logf("response: %+v\n", response)
}
