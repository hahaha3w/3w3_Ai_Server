package usecase

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hahaha3w/3w3_Ai_Server/activity/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/activity/pkg/convert"
	"github.com/hahaha3w/3w3_Ai_Server/activity/pkg/count"
	"github.com/hahaha3w/3w3_Ai_Server/activity/pkg/log"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/activity"
	"time"
)

var _ domain.ActivityUsecase = &ConcreteActivityUsecase{}

type ConcreteActivityUsecase struct {
	repo  domain.ActivityRepo
	cache *redis.Client
}

func NewConcreteActivityUsecase(repo domain.ActivityRepo, cache *redis.Client) *ConcreteActivityUsecase {
	return &ConcreteActivityUsecase{
		repo:  repo,
		cache: cache,
	}
}

const (
	CacheKeyActivityListPrefix   = "activity:list:%d"
	CacheKeyActivityDetailPrefix = "activity:detail:%d"
	CacheExpiration              = time.Hour * 24
	EmptyDataExpiration          = time.Minute * 5
)

type CachedList struct {
	Data  *[]domain.Activity
	Total int32
}

func (c ConcreteActivityUsecase) CreateUserActivity(ctx context.Context, userId int64, relationId int64, activityType, description string) (resp *activity.CreateUserActivityResp, err error) {
	// 创建活动对象
	activityModel := &domain.Activity{
		UserID:      int(userId),
		RelatedID:   int(relationId),
		Type:        activityType,
		Description: description,
		CreatedAt:   time.Now(),
	}

	// 保存活动记录
	if err := c.repo.CreateUserActivity(ctx, activityModel); err != nil {
		return nil, err
	}

	// 清理相关缓存
	patterns := []string{
		fmt.Sprintf(CacheKeyActivityListPrefix+":*", userId),   // 列表缓存
		fmt.Sprintf(CacheKeyActivityDetailPrefix+":*", userId), // 详情缓存
	}

	for _, pattern := range patterns {
		keys, err := c.cache.Keys(ctx, pattern).Result()
		if err != nil {
			continue
		}
		if len(keys) > 0 {
			if err := c.cache.Del(ctx, keys...).Err(); err != nil {
				log.Log().Error("Failed to delete %s caches: %v", pattern, err)
			}
		}
	}

	return &activity.CreateUserActivityResp{
		Success:  true,
		Message:  "",
		Activity: convert.DomainActivityToRPCGenActivity(activityModel),
	}, nil
}

func (c ConcreteActivityUsecase) GetUserActivity(ctx context.Context, userId int64, page, pageSize int32) (resp *activity.GetUserActivityResp, err error) {
	// 规范化缓存键参数
	params := struct {
		Page     int32
		PageSize int32
	}{
		Page:     page,
		PageSize: pageSize,
	}

	cacheKey := fmt.Sprintf("%s:%v",
		fmt.Sprintf(CacheKeyActivityListPrefix, userId),
		md5.Sum([]byte(fmt.Sprintf("%+v", params))), // 使用结构体哈希作为唯一标识
	)

	// 尝试获取完整缓存
	var cached CachedList
	cachedData, err := c.cache.Get(ctx, cacheKey).Bytes()
	if err == nil {
		if err := json.Unmarshal(cachedData, &cached); err == nil {
			if cached.Total > 0 {
				return nil, nil
			}
			return nil, nil
		}
	}

	// 数据库查询
	activities, total, err := c.repo.GetUserActivity(ctx, userId, page, pageSize)
	if err != nil {
		return nil, err
	}

	// 构建缓存对象
	cacheObj := CachedList{
		Data:  activities,
		Total: total,
	}

	expiration := CacheExpiration
	if total == 0 {
		expiration = EmptyDataExpiration
	}

	// 统一存储缓存
	cacheData, err := json.Marshal(cacheObj)
	if err != nil {
		log.Log().Error("Failed to marshal cache data: %v", err)
	} else {
		if err := c.cache.Set(ctx, cacheKey, cacheData, expiration).Err(); err != nil {
			log.Log().Error("Failed to cache activity list: %v", err)
		}
	}

	// 得到用户活动次数统计
	userCount, err := count.GetUserCount(int32(userId), c.cache)
	if err != nil {
		return nil, err
	}

	return &activity.GetUserActivityResp{
		Activities:  convert.DomainActivitiesToRPCGenActivities(activities),
		ChatCount:   int32(userCount.ChatCount),
		MemoirCount: int32(userCount.MemoirCount),
		UseDay:      int32(userCount.UseDays),
	}, nil
}
