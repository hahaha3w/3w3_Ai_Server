package convert

import (
	"github.com/hahaha3w/3w3_Ai_Server/memoir/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/memoir"
)

func DomainMemoirToRPCGenMemoir(domainMemoir *domain.Memoir) *memoir.Memoir {
	if domainMemoir == nil {
		return nil
	}

	// 定义时间格式
	const layout = "2006-01-02 15:04:05"

	// 将 time.Time 转换为字符串
	startDateStr := domainMemoir.StartDate.Format(layout)
	endDateStr := domainMemoir.EndDate.Format(layout)
	createDateStr := domainMemoir.CreatedAt.Format(layout)

	return &memoir.Memoir{
		Id:        int32(domainMemoir.MemoirID),
		UserId:    int32(domainMemoir.UserID),
		Title:     domainMemoir.Title,
		Content:   domainMemoir.Content,
		Type:      domainMemoir.Type,
		Style:     domainMemoir.Style,
		StartDate: startDateStr,
		EndDate:   endDateStr,
		CreatedAt: createDateStr,
	}
}

func DomainMemoirsToRPCGenMemoirs(domainMemoirs []*domain.Memoir) []*memoir.Memoir {
	if len(domainMemoirs) == 0 {
		return nil
	}

	// 定义时间格式
	const layout = "2006-01-02 15:04:05"

	// 创建一个切片用于存储转换后的结果
	rpcGenMemoirs := make([]*memoir.Memoir, 0, len(domainMemoirs))

	// 遍历 domainMemoirs，逐个转换
	for _, domainMemoir := range domainMemoirs {
		if domainMemoir == nil {
			continue // 跳过空值
		}

		// 将 time.Time 转换为字符串
		startDateStr := domainMemoir.StartDate.Format(layout)
		endDateStr := domainMemoir.EndDate.Format(layout)
		createDateStr := domainMemoir.CreatedAt.Format(layout)

		// 转换为 memoir.Memoir
		rpcGenMemoir := &memoir.Memoir{
			Id:        int32(domainMemoir.MemoirID),
			UserId:    int32(domainMemoir.UserID),
			Title:     domainMemoir.Title,
			Content:   domainMemoir.Content,
			Type:      domainMemoir.Type,
			Style:     domainMemoir.Style,
			StartDate: startDateStr,
			EndDate:   endDateStr,
			CreatedAt: createDateStr,
		}

		// 添加到结果切片中
		rpcGenMemoirs = append(rpcGenMemoirs, rpcGenMemoir)
	}

	return rpcGenMemoirs
}
