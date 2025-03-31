package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/infra/rpc"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/activity"
	"strconv"
)

type ActivityApi struct{}

// GetUserActivities 获取用户活动列表
func (api *ActivityApi) GetUserActivities(ctx *gin.Context) {
	var req activity.GetUserActivityReq

	// 从查询参数中获取字符串
	pageStr := ctx.DefaultQuery("page", "1")           // 默认值为 "1"
	pageSizeStr := ctx.DefaultQuery("page_size", "15") // 默认值为 "15"

	// 将字符串转换为 int32
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	// 赋值给 Protobuf 结构体
	req.Page = int32(page)
	req.PageSize = int32(pageSize)

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
	activities := make([]*domain.ActivityResp, 0)
	for _, a := range resp.Activities {
		activities = append(activities, &domain.ActivityResp{
			ActivityID:  a.ActivityId,
			UserID:      a.UserId,
			RelationID:  a.RelationId,
			Type:        a.Type,
			Description: a.Description,
			CreatedAt:   a.CreatedAt,
		})
	}
	vo := domain.ActivityListResp{
		HasMore:     resp.HasMore,
		Total:       resp.Total,
		UseDay:      resp.UseDay,
		ChatCount:   resp.ChatCount,
		MemoirCount: resp.MemoirCount,
		Activities:  activities,
	}
	// 返回成功响应
	domain.Success(ctx, vo)
}
