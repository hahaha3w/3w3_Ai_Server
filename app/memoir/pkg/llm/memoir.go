package llm

import (
	"context"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/hahaha3w/3w3_Ai_Server/common/llm"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/pkg/log"
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

type Content struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func createAnalysisTemplate() prompt.ChatTemplate {
	return prompt.FromMessages(schema.FString,
		schema.SystemMessage("{system}"),
		schema.UserMessage("问题：{questions}"),
	)
}
func createMemoirTemplate() prompt.ChatTemplate {
	return prompt.FromMessages(schema.FString,
		schema.SystemMessage(memoirPrompt),
		schema.UserMessage("问题：{questions}"),
	)
}

func createConverterLambda() *compose.Lambda {
	lambda := compose.InvokableLambda(func(ctx context.Context, input *schema.Message) (key map[string]any, err error) {
		key = map[string]any{}
		key["assistant"] = input.Content
		return key, nil
	})
	return lambda
}

func GenerateMemoir(ctx context.Context, questions string, style string) (title string, body string, err error) {

	workflow := compose.NewGraph[map[string]any, *Content]()
	_ = workflow.AddChatModelNode(ModelNodeAnalysis, llm.GetOpenAIChatModel())
	_ = workflow.AddChatModelNode(ModelNodeMemoir, llm.GetOpenAIChatModel())
	_ = workflow.AddChatTemplateNode(PromptNodeAnalysis, createAnalysisTemplate())
	_ = workflow.AddChatTemplateNode(PromptNodeMemoir, createMemoirTemplate())
	_ = workflow.AddLambdaNode(ConverterLambda1, createConverterLambda())
	_ = workflow.AddLambdaNode(ConverterLambda2, createConverterLambda())
	_ = workflow.AddLambdaNode(ReactLambda, GetReactAgentLambdaNode(ctx))
	//edge 1
	_ = workflow.AddEdge(compose.START, PromptNodeAnalysis)
	_ = workflow.AddEdge(PromptNodeAnalysis, ModelNodeAnalysis)
	_ = workflow.AddEdge(ModelNodeAnalysis, ConverterLambda1)
	_ = workflow.AddEdge(ConverterLambda1, PromptNodeMemoir)
	//edge 2
	_ = workflow.AddEdge(compose.START, PromptNodeMemoir)

	//edgeFinal
	_ = workflow.AddEdge(PromptNodeMemoir, ModelNodeMemoir)
	_ = workflow.AddEdge(ModelNodeMemoir, ReactLambda)
	_ = workflow.AddEdge(ReactLambda, compose.END)

	r, err := workflow.Compile(ctx, compose.WithNodeTriggerMode(compose.AllPredecessor))
	if err != nil {
		log.Log().Error(err)
		return
	}

	in := map[string]any{"questions": questions, "system": analysisPrompt, "range": "周", "style": style}
	ret, err := r.Invoke(ctx, in)
	if err != nil {
		log.Log().Error(err)
		return
	}

	return ret.Title, ret.Body, nil
}
