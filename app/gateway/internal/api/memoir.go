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
	memoirs := make([]domain.Memoir, 0)
	for _, m := range resp.Memoirs {

		memoirs = append(memoirs, domain.Memoir{
			Id:        m.Id,
			UserId:    m.UserId,
			Title:     m.Title,
			Content:   m.Content,
			Type:      m.Type,
			Style:     m.Style,
			EndDate:   m.EndDate,
			CreatedAt: m.CreatedAt,
		})
	}
	vo := domain.ListMemoirResp{
		Memoirs: memoirs,
		Total:   resp.Total,
		HasMore: resp.HasMore,
	}

	// 返回成功响应
	domain.Success(ctx, vo)
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
	vo := domain.Memoir{
		Id:        resp.Memoir.Id,
		UserId:    resp.Memoir.UserId,
		Title:     resp.Memoir.Title,
		Content:   resp.Memoir.Content,
		Type:      resp.Memoir.Type,
		Style:     resp.Memoir.Style,
		EndDate:   resp.Memoir.EndDate,
		CreatedAt: resp.Memoir.CreatedAt,
	}

	// 返回成功响应
	domain.Success(ctx, vo)
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
	vo := domain.GenerateMemoirResp{

		Success: memoirCase.Success,
		Memoir: domain.Memoir{
			Id:        memoirCase.Memoir.Id,
			UserId:    memoirCase.Memoir.UserId,
			Title:     memoirCase.Memoir.Title,
			Content:   memoirCase.Memoir.Content,
			Type:      memoirCase.Memoir.Type,
			Style:     memoirCase.Memoir.Style,
			EndDate:   memoirCase.Memoir.EndDate,
			CreatedAt: memoirCase.Memoir.CreatedAt,
		},
	}
	domain.Success(ctx, vo)

}
