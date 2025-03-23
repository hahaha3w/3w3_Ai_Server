package llm

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
	"github.com/spf13/viper"
	"log"
)

func GetOpenAIChatModel() model.ChatModel {

	c, err := openai.NewChatModel(context.Background(), &openai.ChatModelConfig{
		APIKey:  viper.GetString("llm.token"),
		BaseURL: "https://openrouter.ai/api/v1",
		Model:   "anthropic/claude-3.7-sonnet",
	})
	if err != nil {
		log.Fatal(err)
	}
	return c
}
func Stream(ctx context.Context, llm model.ChatModel, in []*schema.Message) *schema.StreamReader[*schema.Message] {
	stream, err := llm.Stream(ctx, in)
	if err != nil {
		log.Fatal(err)
	}
	return stream
}
