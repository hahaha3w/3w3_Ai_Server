package llm

import (
	"context"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

type ContentInfoRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func GetContentInfoTool() tool.InvokableTool {
	ContentInfoTool := utils.NewTool(
		&schema.ToolInfo{
			Name: "extract_title_and_body",
			Desc: "提取内容中的标题和正文",
			ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
				"title": {
					Type: "string",
					Desc: "内容的标题",
				},
				"body": {
					Type: "string",
					Desc: "内容的正文",
				},
			})},
		func(ctx context.Context, resp ContentInfoRequest) (ContentInfoRequest ContentInfoRequest, err error) {
			if err != nil {
				return resp, err
			}
			return resp, nil
		},
	)
	return ContentInfoTool
}
