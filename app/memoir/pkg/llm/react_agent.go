package llm

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/hahaha3w/3w3_Ai_Server/common/llm"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/pkg/log"
	"time"
)

func GetReactAgentLambdaNode(ctx context.Context) *compose.Lambda {

	rAgent, err := react.NewAgent(ctx, &react.AgentConfig{
		Model: llm.GetOpenAIChatModel(),
		ToolsConfig: compose.ToolsNodeConfig{
			Tools: []tool.BaseTool{GetContentInfoTool()},
		},
		MaxStep: 10,
	})
	if err != nil {
		log.Log().Errorf("failed to create agent: %v", err)
		return nil
	}
	agentLambda := compose.InvokableLambda(func(ctx context.Context, input *schema.Message) (output *Content, err error) {
		p, err := prompt.FromMessages(schema.FString,
			schema.SystemMessage("{system}"),
			schema.UserMessage("{content}")).Format(ctx, map[string]any{
			"content": input.Content,
			"system":  convertJsonTemplate,
		})
		if err != nil {
			log.Log().Errorf("failed to parse message: %v", err)
			return nil, err
		}
		ret, err := rAgent.Generate(ctx, p)
		if err != nil {
			log.Log().Errorf("failed to generate message: %v", err)
			return nil, err
		}
		content := &Content{}
		err = json.Unmarshal([]byte(ret.Content), content)
		if err != nil {
			log.Log().Errorf("failed to parse message: %v", err)
			content.Title = time.Now().String()
			content.Body = ret.Content
			return content, nil
		}
		return content, nil

	})

	return agentLambda

}
