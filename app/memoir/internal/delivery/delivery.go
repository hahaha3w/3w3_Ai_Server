package delivery

import (
	"context"
	"github.com/google/wire"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/pkg/convert"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/pkg/log"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/memoir"
)

var ProviderSet = wire.NewSet(New)

type MemoirDelivery struct {
	repo    domain.MemoirRepo
	usecase domain.MemoirUsecase
}

func New(repo domain.MemoirRepo, usecase domain.MemoirUsecase) *MemoirDelivery {
	return &MemoirDelivery{
		repo:    repo,
		usecase: usecase,
	}
}

// GenerateMemoir 生成回忆录
func (d *MemoirDelivery) GenerateMemoir(ctx context.Context, req *memoir.GenerateMemoirRequest) (res *memoir.GenerateMemoirResponse, err error) {
	// 调用 usecase 层的 GenerateMemoir 方法
	generatedMemoir, err := d.usecase.GenerateMemoir(ctx, int(req.UserId), req.Style, req.Type, req.StartDate, req.EndDate)
	if err != nil {
		log.Log().Error(err)
		return &memoir.GenerateMemoirResponse{
			Success:  false,
			Memoir:   nil,
			ErrorMsg: err.Error(),
		}, err
	}

	// 封装响应
	return &memoir.GenerateMemoirResponse{
		Success:  true,
		Memoir:   convert.DomainMemoirToRPCGenMemoir(generatedMemoir),
		ErrorMsg: "",
	}, nil
}

// GetMemoirList 获取回忆录列表
func (d *MemoirDelivery) GetMemoirList(ctx context.Context, req *memoir.GetMemoirListRequest) (res *memoir.GetMemoirListResponse, err error) {
	// 调用 usecase 层的 GetMemoirList 方法
	memoirs, hasMore, err := d.usecase.GetMemoirList(ctx, int(req.UserId), req.Type, req.StartDate, req.EndDate, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 封装响应
	return &memoir.GetMemoirListResponse{
		Memoirs: convert.DomainMemoirsToRPCGenMemoirs(memoirs),
		Total:   int32(len(memoirs)),
		HasMore: hasMore,
	}, nil
}

// GetMemoirDetail 获取回忆录详情
func (d *MemoirDelivery) GetMemoirDetail(ctx context.Context, req *memoir.GetMemoirDetailRequest) (res *memoir.GetMemoirDetailResponse, err error) {
	// 调用 usecase 层的 GetMemoirDetail 方法
	memoirDetail, err := d.usecase.GetMemoirDetail(ctx, int(req.MemoirId), int(req.UserId))
	if err != nil {
		return nil, err
	}

	// 封装响应
	res = &memoir.GetMemoirDetailResponse{
		Memoir: convert.DomainMemoirToRPCGenMemoir(memoirDetail),
	}
	return res, nil
}

// DeleteMemoir 删除回忆录
func (d *MemoirDelivery) DeleteMemoir(ctx context.Context, req *memoir.DeleteMemoirRequest) (res *memoir.DeleteMemoirResponse, err error) {
	// 调用 usecase 层的 DeleteMemoir 方法
	if err := d.usecase.DeleteMemoir(ctx, int(req.MemoirId), int(req.UserId)); err != nil {
		return &memoir.DeleteMemoirResponse{
			Success:  false,
			ErrorMsg: err.Error(),
		}, err
	}

	// 封装响应
	return &memoir.DeleteMemoirResponse{
		Success:  true,
		ErrorMsg: "",
	}, nil
}
