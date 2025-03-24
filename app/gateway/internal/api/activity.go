package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/infra/rpc"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/activity"
)

type ActivityApi struct{}

// GetUserActivities 获取用户活动列表
func (api *ActivityApi) GetUserActivities(ctx *gin.Context) {
	var req activity.GetUserActivityReq

	// 使用 ShouldBindQuery 绑定 Query 参数
	if err := ctx.ShouldBindQuery(&req); err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	// 设置 userId
	userId, err := domain.GetUserIdFromContext(ctx)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}
	req.UserId = int64(userId)

	// 调用 RPC 客户端获取用户活动列表
	resp, err := rpc.ActivityClient.GetUserActivities(ctx, &req)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	// 返回成功响应
	domain.Success(ctx, resp)
}

// CreateUserActivity 创建用户活动
func (api *ActivityApi) CreateUserActivity(ctx *gin.Context) {
	var req activity.CreateUserActivityReq

	// 使用 ShouldBindJSON 绑定 JSON 请求体
	if err := ctx.ShouldBindJSON(&req); err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	// 设置 userId
	userId, err := domain.GetUserIdFromContext(ctx)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}
	req.UserId = int64(userId)

	// 调用 RPC 客户端创建用户活动
	resp, err := rpc.ActivityClient.CreateUserActivity(ctx, &req)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	// 返回成功响应
	domain.Success(ctx, resp)
}
