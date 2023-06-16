package chat

import "testing"

const (
	apiKey  = "1efd12c080a64d8a997bfead8e061938"
	baseURL = "https://sys-devops.openai.azure.com/"
)

func TestCreateChatCompletionWithContent(t *testing.T) {
	g := NewGPT(apiKey, baseURL, WithModel("Sys-Devops"))
	response, err := g.SimpleChatCompletion("hello")
	if err != nil {
		t.FailNow()
	}
	t.Logf("response: %+v", response.Context)
}
