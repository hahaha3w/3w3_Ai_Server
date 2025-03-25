package llm

import (
	"context"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"log"
)

func CreateTemplate() prompt.ChatTemplate {
	return prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。"),
		schema.SystemMessage("输出格式为markdown格式"),
		schema.MessagesPlaceholder("chat_history", false),
		schema.UserMessage("问题：{question}"),
	)
}

func CreateMessagesFromTemplate(question string) []*schema.Message {
	template := CreateTemplate()
	values := map[string]interface{}{
		"role":         "ai助手",
		"style":        "专业",
		"question":     question,
		"chat_history": []*schema.Message{}}
	messages, err := template.Format(context.Background(), values)
	if err != nil {
		log.Fatal(err)
	}

	return messages
}
