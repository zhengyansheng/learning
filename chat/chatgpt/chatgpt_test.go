package chatgpt

import (
	"testing"

	"github.com/sashabaranov/go-openai"
)

var (
	sqlOptimization = "优化建议：1. 当只要一行数据时使用 LIMIT 1, 2. 为搜索字段建索引, 3. 在Join表的时候使用相当类型的例，并将其索引, 4. 千万不要 ORDER BY RAND(), 5. 避免 SELECT *, 6. 永远为每张表设置一个ID, 7. 使用 ENUM 而不是 VARCHAR, 8. 对查询进行优化，要尽量避免全表扫描, 9. 应尽量避免在 where 子句中对字段进行 null 值判断, 10. 应尽量避免在 where 子句中使用 != 或 <> 操作符, 11. 应尽量避免在 where 子句中使用 or 来连接条件，否则系统将进行全表扫描, 12. 模糊查询前边不要使用%，否则系统将可能无法正确使用索引, 13. 应尽量避免在where子句中对字段进行函数操作，否则系统将可能无法正确使用索引, 14. 不要在 where 子句中的=左边进行函数、算术运算或其他表达式运算，否则系统将可能无法正确使用索引, 15. 不带条件的查询是全表扫描，不允许发起, 16. 尽量避免大事务操作，提高系统并发能力"

	sqlContent = "按照上面的优化建议，优化这条sql语句并提出建议: select id, code, `type`, biz_type, category, user_id, student_id, user_shipping_address_id, country_id, country_name, province_id, province_name, city_id, city_name, area_id, area_name, address, receive_contact_name, receive_contact_phone, receive_phone_code, receive_phone_country_code, pay_account_id, amounts_payable, amounts_actuallypay, create_time, order_status, pay_status, waybill_number, express_order_id, logistics_company, send_time, business_region, business_type, ware_house_country_id from mall_order WHERE ( `type` not in ( 2 , 3 ) and create_time <= '2023-03-12 00:05:00.008' and order_status = 0 )"
)

func TestChatGPTSimple(t *testing.T) {
	apiAuthTokens := []string{
		"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",
	}
	g := NewGPT(apiAuthTokens, WithModel(openai.GPT3Dot5Turbo0301), WithTimeout(60))
	response, err := g.SimpleRequest(sqlContent)
	if err != nil {
		panic(err)
	}
	t.Logf("response: %+v\n", response.Reply)
}

func TestChatGPTRequest(t *testing.T) {
	apiAuthTokens := []string{
		"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",
	}
	var newMessages = []openai.ChatCompletionMessage{
		DefaultSystemChatCompletionMessage(),
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
			Content: sqlOptimization,
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "我已经学会了SQL优化标准，接下来SQL语句我将按照以上标准执行优化, 建议和提供优化后的SQL",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: sqlContent,
		},
	}

	g := NewGPT(apiAuthTokens, WithModel(openai.GPT3Dot5Turbo0301), WithTimeout(60))

	response, err := g.Request(newMessages)
	if err != nil {
		panic(err)
	}
	t.Logf("response: %+v\n", response)
}
