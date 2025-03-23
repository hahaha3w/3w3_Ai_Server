package llm

import (
	"context"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"log"
)

func CreateTemplate() prompt.ChatTemplate {
	return prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。你的目标是帮助程序员保持积极乐观的心态，提供技术建议的同时也要关注他们的心理健康。"),
		schema.MessagesPlaceholder("chat_history", false),
		schema.UserMessage("问题：{question}"),
	)
}

func CreateMessagesFromTemplate(question string) []*schema.Message {
	template := CreateTemplate()
	values := map[string]interface{}{
		"role":     "程序员鼓励师",
		"style":    "像王宏亮一样温柔",
		"question": question,
		"chat_history": []*schema.Message{
			schema.UserMessage("你好"),
			schema.AssistantMessage("你好，我是程序员鼓励师。我可以帮助你保持积极乐观的心态，提供技术建议的同时也要关注你的心理健康。", nil),
		},
	}
	messages, err := template.Format(context.Background(), values)
	if err != nil {
		log.Fatal(err)
	}

	return messages
}
