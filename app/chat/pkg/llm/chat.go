package llm

import (
	"context"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/hahaha3w/3w3_Ai_Server/chat/pkg/log"
	"github.com/hahaha3w/3w3_Ai_Server/common/llm"
)

const (
	ModelNodeAnalysis = "model_analysis" // 分析模型节点
	ModelNodeMemoir   = "model_memoir"   // 传记模型节点

	PromptNodeAnalysis    = "prompt_analysis" // 系统提示模板
	PromptNodeMemoir      = "prompt_memoir"   // 传记提示模板
	PromptNodeJsonConvert = "prompt_json_convert"

	ConverterLambda1 = "converter_lambda1" // 格式转换器
	ConverterLambda2 = "converter_lambda2"
	ReactLambda      = "react_lambda"
)

func getChatPromptTemplate() prompt.ChatTemplate {
	return prompt.FromMessages(schema.FString,
		schema.SystemMessage(chatPrompt),
		schema.MessagesPlaceholder("chat_histories", false),
		schema.UserMessage("{question}"),
	)

}

func Chat(ctx context.Context, question string, mood string) (*schema.StreamReader[*schema.Message], error) {

	workflow := compose.NewChain[map[string]any, *schema.Message]()
	workflow.
		AppendChatTemplate(getChatPromptTemplate()).
		AppendChatModel(llm.GetOpenAIChatModel())

	r, err := workflow.Compile(ctx)
	if err != nil {
		log.Log().Error(err)
		return nil, err
	}

	in := map[string]any{
		"question":       question,
		"mood":           mood,
		"chat_histories": []*schema.Message{},
	}
	ret, err := r.Stream(ctx, in)
	if err != nil {
		log.Log().Error(err)
		return nil, err
	}

	return ret, nil
}
