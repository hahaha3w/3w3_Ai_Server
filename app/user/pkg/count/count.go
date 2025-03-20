package count

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hahaha3w/3w3_Ai_Server/user/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/user/pkg/log"
	"time"
)

// InitUserCount 初始化用户统计数据
func InitUserCount(userId int32, cache *redis.Client) (err error) {
	key := fmt.Sprintf("user:count:%d", userId)

	// 初始化默认值
	now := time.Now()
	userCount := map[string]interface{}{
		"chat_count":   0,
		"memoir_count": 0,
		"use_days":     1,
		"last_updated": now.Format(time.RFC3339),
	}

	// 将数据存储到 Redis 哈希中
	err = cache.HSet(context.Background(), key, userCount).Err()
	if err != nil {
		log.Log().Error(err)
		return errors.New("failed to initialize user count in Redis")
	}

	return nil
}

// UpdateUserCount 更新用户统计数据
func UpdateUserCount(userId int64, userCount *domain.UpdateUserCount, cache *redis.Client) (err error) {
	key := fmt.Sprintf("user:count:%d", userId)

	// 更新数据
	err = cache.HIncrBy(context.Background(), key, "chat_count", userCount.ChatCountIncr).Err()
	if err != nil {
		return errors.New("failed to update chat count in Redis")
	}

	err = cache.HIncrBy(context.Background(), key, "memoir_count", userCount.MemoirCountIncr).Err()
	if err != nil {
		return errors.New("failed to update memoir count in Redis")
	}

	// 判断是否超过一天
	lastUpdated, err := cache.HGet(context.Background(), key, "last_updated").Result()
	if err != nil {
		return fmt.Errorf("failed to get last updated time from Redis: %w", err)
	}

	// 解析 last_updated
	lastTime, err := time.Parse(time.RFC3339, lastUpdated)
	if err != nil {
		return fmt.Errorf("failed to parse last updated time: %w", err)
	}

	// 判断上次更新时间是否超过一天
	todayTime := time.Now()
	if todayTime.Sub(lastTime) >= 24*time.Hour && userCount.UseDaysIncr != 0 {
		// 增加 use_days
		err = cache.HIncrBy(context.Background(), key, "use_days", userCount.UseDaysIncr).Err()
		if err != nil {
			return fmt.Errorf("failed to increment use_days: %w", err)
		}
	}

	return nil
}
