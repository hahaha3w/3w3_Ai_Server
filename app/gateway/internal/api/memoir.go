package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/infra/rpc"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/memoir"
	"strconv"
)

type MemoirApi struct {
}

// GenerateMemoir 生成回忆录
func (api *MemoirApi) GenerateMemoir(ctx *gin.Context) {
	var req memoir.GenerateMemoirRequest
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
	req.UserId = userId

	resp, err := rpc.MemoirClient.GenerateMemoir(ctx, &req)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	domain.Success(ctx, resp)
}

// GetMemoirList 获取已生成的回忆录列表
func (api *MemoirApi) GetMemoirList(ctx *gin.Context) {
	var req memoir.GetMemoirListRequest

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
	req.UserId = userId

	// 调用 RPC 客户端获取回忆录列表
	resp, err := rpc.MemoirClient.GetMemoirList(ctx, &req)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	// 返回成功响应
	domain.Success(ctx, resp)
}

// GetMemoirDetail 获取回忆录详情
func (api *MemoirApi) GetMemoirDetail(ctx *gin.Context) {
	var req memoir.GetMemoirDetailRequest

	// 从 URL 路径中获取 memoir_id
	memoirIdStr := ctx.Param("memoir_id")
	if memoirIdStr == "" {
		domain.ErrorMsg(ctx, "memoir_id is required")
		return
	}
	memoirId, err := strconv.ParseInt(memoirIdStr, 10, 64)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}
	req.MemoirId = int32(memoirId)

	// 设置 userId
	userId, err := domain.GetUserIdFromContext(ctx)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}
	req.UserId = userId

	resp, err := rpc.MemoirClient.GetMemoirDetail(ctx, &req)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	// 返回成功响应
	domain.Success(ctx, resp)
}

// DeleteMemoir 删除回忆录
func (api *MemoirApi) DeleteMemoir(ctx *gin.Context) {
	var req memoir.DeleteMemoirRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	// 从 URL 路径中获取 memoir_id
	memoirIdStr := ctx.Param("memoir_id")
	if memoirIdStr == "" {
		domain.ErrorMsg(ctx, "memoir_id is required")
		return
	}
	memoirId, err := strconv.ParseInt(memoirIdStr, 10, 64)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}
	req.MemoirId = int32(memoirId)

	// 设置 userId
	userId, err := domain.GetUserIdFromContext(ctx)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}
	req.UserId = userId

	resp, err := rpc.MemoirClient.DeleteMemoir(ctx, &req)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	domain.Success(ctx, resp)
}
