package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/infra/rpc"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/memoir"
	"strconv"
)

type MemoirApi struct {
}

// GetMemoirList 获取已生成的回忆录列表
func (api *MemoirApi) GetMemoirList(ctx *gin.Context) {
	var req memoir.GetMemoirListRequest

	// 从查询参数中获取字符串并设置默认值
	pageStr := ctx.DefaultQuery("page", "1")           // 默认值为 "1"
	pageSizeStr := ctx.DefaultQuery("page_size", "15") // 默认值为 "15"
	typeStr := ctx.Query("type")                       // 可选参数
	styleStr := ctx.Query("style")                     // 可选参数
	startDateStr := ctx.Query("start_date")            // 可选参数
	endDateStr := ctx.Query("end_date")                // 可选参数

	// 将字符串转换为 int32
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	// 赋值给 Protobuf 结构体
	req.Page = int32(page)
	req.PageSize = int32(pageSize)

	// 处理其他可选参数
	if typeStr != "" {
		req.Type = typeStr
	}
	if styleStr != "" {
		req.Style = styleStr
	}
	if startDateStr != "" {
		req.StartDate = startDateStr
	}
	if endDateStr != "" {
		req.EndDate = endDateStr
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
func (api *MemoirApi) GenerateMemoir(ctx *gin.Context) {
	var req memoir.GenerateMemoirRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err)
		domain.ErrorMsg(ctx, err.Error())
		return
	}
	userId, err := domain.GetUserIdFromContext(ctx)
	if err != nil {
		fmt.Println(err)
		domain.ErrorMsg(ctx, err.Error())
		return
	}
	req.UserId = userId
	memoirCase, err := rpc.MemoirClient.GenerateMemoir(ctx, &req)
	if err != nil {
		fmt.Println(err)
		domain.ErrorMsg(ctx, err.Error())
		return
	}
	if memoirCase == nil || memoirCase.Success == false {
		domain.ErrorMsg(ctx, "memoir generation failed")
		return
	}
	domain.Success(ctx, memoirCase)

}
