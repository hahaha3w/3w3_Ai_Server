package convert

import (
	"github.com/hahaha3w/3w3_Ai_Server/activity/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/activity"
)

func DomainActivityToRPCGenActivity(domainActivity *domain.Activity) *activity.Activity {
	if domainActivity == nil {
		return nil
	}

	// 将 time.Time 转换为字符串
	createdAtStr := domainActivity.CreatedAt.Format("2006-01-02 15:04:05")

	return &activity.Activity{
		ActivityId:  int64(domainActivity.ActivityID),
		UserId:      int64(domainActivity.UserID),
		RelationId:  int64(int32(domainActivity.RelatedID)),
		Type:        domainActivity.Type,
		Description: domainActivity.Description,
		CreatedAt:   createdAtStr,
	}
}

func DomainActivitiesToRPCGenActivities(domainActivities []*domain.Activity) []*activity.Activity {
	if domainActivities == nil {
		return nil
	}

	var rpcGenActivities []*activity.Activity
	for _, domainActivity := range domainActivities {
		rpcGenActivities = append(rpcGenActivities, DomainActivityToRPCGenActivity(domainActivity))
	}

	return rpcGenActivities
}
