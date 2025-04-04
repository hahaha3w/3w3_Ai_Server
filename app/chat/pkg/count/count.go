package count

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/chat/pkg/log"
	"strconv"
	"time"
)

// GetUserCount 获取用户统计数据
func GetUserCount(userId int32, cache *redis.Client) (*domain.UserCount, error) {
	key := fmt.Sprintf("user:count:%d", userId)

	// 从 Redis 哈希中获取数据
	result, err := cache.HGetAll(context.Background(), key).Result()
	if err != nil {
		log.Log().Error(err)
		return nil, errors.New("failed to get user count from Redis")
	}

	// 解析 map 到 UserCount 结构体
	userCount := &domain.UserCount{}
	if chatCount, ok := result["chat_count"]; ok {
		userCount.ChatCount, _ = strconv.ParseInt(chatCount, 10, 64)
	}
	if memoirCount, ok := result["memoir_count"]; ok {
		userCount.MemoirCount, _ = strconv.ParseInt(memoirCount, 10, 64)
	}
	if useDays, ok := result["use_days"]; ok {
		userCount.UseDays, _ = strconv.ParseInt(useDays, 10, 64)
	}

	// 获取并更新上次更新时间
	lastUpdated, err := cache.HGet(context.Background(), key, "last_updated").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get last updated time from Redis: %w", err)
	}

	// 更新上次更新时间
	now := time.Now()
	err = cache.HSet(context.Background(), key, "last_updated", now.Format(time.RFC3339)).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to update last updated time in Redis: %w", err)
	}

	// 判断是否同一天登录
	if lastUpdated != "" {
		lastTime, err := time.Parse(time.RFC3339, lastUpdated)
		if err != nil {
			return nil, fmt.Errorf("failed to parse last updated time: %w", err)
		}

		// 判断是否跨自然天（考虑时区）
		lastDate := lastTime.UTC().Truncate(24 * time.Hour)
		todayDate := now.UTC().Truncate(24 * time.Hour)

		if lastDate.Before(todayDate) {
			// 自然天数增加（每天首次触发）
			incrBy := cache.HIncrBy(context.Background(), key, "use_days", 1)
			if incrBy.Err() != nil {
				return nil, errors.New("failed to update use days in Redis")
			}
			userCount.UseDays = incrBy.Val()
		}
	}

	return userCount, nil
}

// InitUserCount 初始化用户统计数据
func InitUserCount(userId int32, cache *redis.Client) error {
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
	err := cache.HSet(context.Background(), key, userCount).Err()
	if err != nil {
		log.Log().Error(err)
		return errors.New("failed to initialize user count in Redis")
	}

	return nil
}

// UpdateUserCount 更新用户统计数据
func UpdateUserCount(userUpdateCount *domain.UpdateUserCount, cache *redis.Client) (userCount *domain.UserCount, err error) {
	key := fmt.Sprintf("user:count:%d", userUpdateCount.UserId)

	userCount = &domain.UserCount{}

	// 更新数据
	incrBy := cache.HIncrBy(context.Background(), key, "chat_count", userUpdateCount.ChatCountIncr)
	if incrBy.Err() != nil {
		return nil, errors.New("failed to update chat count in Redis")
	}
	userCount.ChatCount = incrBy.Val()

	incrBy = cache.HIncrBy(context.Background(), key, "memoir_count", userUpdateCount.MemoirCountIncr)
	if incrBy.Err() != nil {
		return nil, errors.New("failed to update memoir count in Redis")
	}
	userCount.MemoirCount = incrBy.Val()

	// 判断是否超过一天并更新最新时间
	lastUpdated, err := cache.HGet(context.Background(), key, "last_updated").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get last updated time from Redis: %w", err)
	}
	err = cache.HSet(context.Background(), key, "last_updated", time.Now().Format(time.RFC3339)).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to update last updated time in Redis: %w", err)
	}

	// 解析 last_updated
	lastTime, err := time.Parse(time.RFC3339, lastUpdated)
	if err != nil {
		return nil, fmt.Errorf("failed to parse last updated time: %w", err)
	}

	// 判断上次更新时间是否超过一天
	todayTime := time.Now()
	day := cache.HGet(context.Background(), key, "use_days")
	num, err := strconv.Atoi(day.Val())
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}

	// 将 int 转换为 int64
	num64 := int64(num)
	userCount.UseDays = num64
	// 判断是否跨自然天（考虑时区）
	lastDate := lastTime.UTC().Truncate(24 * time.Hour)
	todayDate := todayTime.UTC().Truncate(24 * time.Hour)

	if lastDate.Before(todayDate) && userUpdateCount.UseDaysIncr != 0 {
		// 自然天数增加（每天首次触发）
		incrBy := cache.HIncrBy(context.Background(), key, "use_days", userUpdateCount.UseDaysIncr)
		if incrBy.Err() != nil {
			return nil, errors.New("failed to update use days in Redis")
		}
		userCount.UseDays = incrBy.Val()
	}

	return userCount, nil
}

// DeleteUserCount 删除用户统计数据
func DeleteUserCount(userId int32, cache *redis.Client) error {
	key := fmt.Sprintf("user:count:%d", userId)

	// 删除 Redis 中的哈希键
	err := cache.Del(context.Background(), key).Err()
	if err != nil {
		log.Log().Error(err)
		return errors.New("failed to delete user count from Redis")
	}

	return nil
}
