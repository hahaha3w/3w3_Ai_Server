package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"sort"
	"time"

	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/chat/pkg/log"
)

const (
	ConvMessagesKey = "conv:%d:messages"
)

func (c *ChatRepoWithCache) StoreMessageRecord(ctx context.Context, m *domain.Message) (userID int, err error) {

	// 构造Redis key
	cacheKey := fmt.Sprintf(ConvMessagesKey, m.ConversationID)

	// 将消息序列化为JSON
	messageJSON, err := json.Marshal(m)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal message: %v", err)
	}

	// 使用Pipeline批量更新缓存
	pipe := c.redis.Pipeline()
	// 存储消息
	pipe.HSet(ctx, cacheKey, fmt.Sprintf("msg:%d", m.MessageID), string(messageJSON))
	// 设置过期时间为1小时
	pipe.Expire(ctx, cacheKey, time.Hour)
	// 更新会话的score
	err = c.mq.PublishConversation(m.ConversationID, m.SendTime)
	if err != nil {
		log.Log().Error(err)
		return userID, err
	}

	_, err = pipe.Exec(ctx)
	//缓存失败，不影响结果
	if err != nil {
		log.Log().Info(err)
		return m.UserID, nil
	}

	return m.UserID, nil
}

func (c *ChatRepoWithCache) ListMessages(ctx context.Context, conversationID int, start, end int) ([]*domain.Message, error) {
	// 构造Redis key
	cacheKey := fmt.Sprintf(ConvMessagesKey, conversationID)

	// 使用HGETALL一次性获取所有消息
	result, err := c.redis.HGetAll(ctx, cacheKey).Result()
	if err != nil {
		log.Log().Info(err)
		// 缓存获取失败时直接从数据库获取
		return c.repo.ListMessage(ctx, conversationID, start, end)
	}

	// 将消息转换为结构体并记录ID
	messages := make([]*domain.Message, 0, len(result))
	for _, val := range result {
		message := &domain.Message{}
		if err := json.Unmarshal([]byte(val), message); err != nil {
			continue
		}
		messages = append(messages, message)
	}

	// 按消息ID倒序排序
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].MessageID > messages[j].MessageID
	})

	// 处理分页范围
	if start < 0 {
		start = 0
	}
	fmt.Println(start, end+1)

	// 如果缓存中的消息数量不足，从数据库获取补充
	if end >= len(messages) || start >= len(messages) {
		// 从数据库获取消息
		dbMessages, err := c.repo.ListMessage(ctx, conversationID, start, end)
		if err != nil {
			log.Log().Error(err)
			return nil, err
		}
		if len(dbMessages) == len(messages) {
			return messages, nil
		}
		// 更新缓存
		pipe := c.redis.Pipeline()
		for _, msg := range dbMessages {
			msgJSON, err := json.Marshal(msg)
			if err != nil {
				log.Log().Error(err)
				continue
			}
			pipe.HSet(ctx, cacheKey, fmt.Sprintf("msg:%d", msg.MessageID), string(msgJSON))
		}
		pipe.Expire(ctx, cacheKey, time.Hour)
		_, err = pipe.Exec(ctx)
		if err != nil {
			log.Log().Info(err)
		}
		log.Log().Info(dbMessages)
		return dbMessages, nil
	}

	// 如果缓存中的消息数量足够，直接返回指定范围的消息
	if end >= len(messages) {
		end = len(messages) - 1
	}
	log.Log().Info(len(messages))
	fmt.Println(start, end+1)

	return messages[start : end+1], nil
}
