package chatgpt

import (
	"testing"

	"github.com/sashabaranov/go-openai"
)

var (
	sqlContent = "优化sql语句并提出建议: select id, code, `type`, biz_type, category, user_id, student_id, user_shipping_address_id, country_id, country_name, province_id, province_name, city_id, city_name, area_id, area_name, address, receive_contact_name, receive_contact_phone, receive_phone_code, receive_phone_country_code, pay_account_id, amounts_payable, amounts_actuallypay, create_time, order_status, pay_status, waybill_number, express_order_id, logistics_company, send_time, business_region, business_type, ware_house_country_id from mall_order WHERE ( `type` not in ( 2 , 3 ) and create_time <= '2023-03-12 00:05:00.008' and order_status = 0 )"
)

func TestChatGPTSimple(t *testing.T) {
	apiAuthTokens := []string{
		"",
	}
	g := NewGPT(apiAuthTokens, WithModel(openai.GPT3Dot5Turbo0301), WithTimeout(60))
	response, err := g.SimpleRequest(sqlContent)
	if err != nil {
		panic(err)
	}
	t.Logf("response: %+v\n", response)
}

func TestChatGPTRequest(t *testing.T) {
	apiAuthTokens := []string{
		"",
	}
	var newMessages = []openai.ChatCompletionMessage{
		DefaultSystemChatCompletionMessage(),
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "hello",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "hello",
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "Hi 亲爱的，有什么问题需要我回答吗？",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "golang 实现快速排序",
		},
	}

	g := NewGPT(apiAuthTokens, WithModel(openai.GPT3Dot5Turbo0301), WithTimeout(60))

	response, err := g.Request(newMessages)
	if err != nil {
		panic(err)
	}
	t.Logf("response: %+v\n", response)
}
