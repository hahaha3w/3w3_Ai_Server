package llm

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/hahaha3w/3w3_Ai_Server/common/llm"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/pkg/log"
)

const (
	ModelNodeAnalysis = "model_analysis" // 分析模型节点
	ModelNodeMemoir   = "model_memoir"   // 传记模型节点

	PromptNodeAnalysis = "prompt_analysis" // 系统提示模板
	PromptNodeMemoir   = "prompt_memoir"   // 传记提示模板

	LambdaConverter = "lambda_converter" // 格式转换器
)

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
func GenerateMemoir(ctx context.Context, questions string, style string) (title string, content string, err error) {

	workflow := compose.NewGraph[map[string]any, *schema.Message]()
	_ = workflow.AddChatModelNode(ModelNodeAnalysis, llm.GetOpenAIChatModel())
	_ = workflow.AddChatModelNode(ModelNodeMemoir, llm.GetOpenAIChatModel())
	_ = workflow.AddChatTemplateNode(PromptNodeAnalysis, createAnalysisTemplate())
	_ = workflow.AddChatTemplateNode(PromptNodeMemoir, createMemoirTemplate())
	_ = workflow.AddLambdaNode(LambdaConverter, createConverterLambda())
	//edge 1
	_ = workflow.AddEdge(compose.START, PromptNodeAnalysis)
	_ = workflow.AddEdge(PromptNodeAnalysis, ModelNodeAnalysis)
	_ = workflow.AddEdge(ModelNodeAnalysis, LambdaConverter)
	_ = workflow.AddEdge(LambdaConverter, PromptNodeMemoir)
	//edge 2
	_ = workflow.AddEdge(compose.START, PromptNodeMemoir)

	//edgeFinal
	_ = workflow.AddEdge(PromptNodeMemoir, ModelNodeMemoir)
	_ = workflow.AddEdge(ModelNodeMemoir, compose.END)

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
	fmt.Println("invoke result: ", ret)
	return "test", ret.Content, nil
}
