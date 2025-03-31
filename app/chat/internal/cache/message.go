package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/chat/pkg/log"
	"strconv"
	"time"
)

const (
	ConvMessagesKey     = "cid:%d:mid:%d"
	ConvMessagesRankKey = "conv:%d:messages:rank"
)

func (c *ChatRepoWithCache) StoreMessageRecord(ctx context.Context, m *domain.Message) (userID int, err error) {
	_, err = c.repo.StoreMessageRecord(ctx, m)
	if err != nil {
		log.Log().Error(err)
		return -1, err
	}
	return m.MessageID, nil
}
func (c *ChatRepoWithCache) ListMessages(ctx context.Context, conversationID int, pageNum, pageSize int) ([]*domain.Message, error) {
	// 计算分页
	start := (pageNum - 1) * pageSize
	end := pageNum*pageSize - 1

	log.Log().Info("尝试从缓存获取消息",
		"conversationID", conversationID,
		"pageNum", pageNum,
		"pageSize", pageSize,
		"start", start,
		"end", end)

	// 1. 从有序集合中倒序获取指定范围的消息ID
	rankKey := fmt.Sprintf(ConvMessagesRankKey, conversationID)
	messageIDs, err := c.redis.ZRevRange(ctx, rankKey, int64(start), int64(end)).Result()

	if err != nil {
		log.Log().Error("Redis ZRevRange失败", "error", err)
	}

	// 2. 如果缓存未命中或不完整，从数据库获取
	if err != nil || len(messageIDs) < end-start+1 {
		log.Log().Info("缓存未命中或不完整，从数据库获取",
			"缓存错误", err,
			"获取到的消息ID数量", len(messageIDs),
			"需要的消息数量", end-start+1)

		// 从数据库获取消息
		messages, err := c.repo.ListMessages(ctx, conversationID, start, end)
		if err != nil {
			log.Log().Error("从数据库获取消息失败", "error", err)
			return nil, fmt.Errorf("从数据库获取消息失败: %w", err)
		}

		log.Log().Info("从数据库成功获取消息", "消息数量", len(messages))

		// 如果有结果，更新缓存
		if len(messages) > 0 {
			c.updateCache(ctx, conversationID, messages)
		}

		return messages, nil
	}

	// 3. 从缓存获取这些消息的具体内容
	pipe := c.redis.Pipeline()
	for _, idStr := range messageIDs {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Log().Error("消息ID格式错误", "idStr", idStr, "error", err)
		}
		msgKey := fmt.Sprintf(ConvMessagesKey, conversationID, id)
		pipe.Get(ctx, msgKey)
	}

	cmds, err := pipe.Exec(ctx)
	if err != nil {
		// Redis错误，回退到数据库
		log.Log().Error("Redis批量获取消息失败，回退到数据库", "error", err)
		return c.repo.ListMessages(ctx, conversationID, start, end)
	}

	// 4. 处理缓存结果
	messages := make([]*domain.Message, 0, len(messageIDs))

	for i, cmd := range cmds {
		if cmd.Err() != nil {
			err = cmd.Err()
			log.Log().Error("处理Redis命令结果出错", "索引", i, "error", cmd.Err())
			break
		}

		msgJSON, _ := cmd.(*redis.StringCmd).Result()

		msg := &domain.Message{}
		if err = json.Unmarshal([]byte(msgJSON), msg); err != nil {
			log.Log().Error("解析消息JSON失败", "error", err)
			break
		}

		messages = append(messages, msg)
	}

	// 5. 如果有消息缺失，从数据库重新获取
	if err != nil {
		log.Log().Info("处理缓存数据过程中出错，回退到数据库", "error", err)

		messages, err := c.repo.ListMessages(ctx, conversationID, start, end)
		if err != nil {
			log.Log().Error("从数据库获取消息失败", "error", err)
			return nil, fmt.Errorf("从数据库获取消息失败: %w", err)
		}

		log.Log().Info("从数据库成功获取消息", "消息数量", len(messages))

		// 更新缓存
		if len(messages) > 0 {
			c.updateCache(ctx, conversationID, messages)
		}

		return messages, nil
	}

	log.Log().Info("成功从缓存获取消息", "消息数量", len(messages))
	return messages, nil
}

// 辅助函数: 更新缓存
func (c *ChatRepoWithCache) updateCache(ctx context.Context, conversationID int, messages []*domain.Message) {

	pipe := c.redis.Pipeline()
	rankKey := fmt.Sprintf(ConvMessagesRankKey, conversationID)

	for _, msg := range messages {
		// 1. 缓存消息内容
		msgJSON, err := json.Marshal(msg)
		if err != nil {
			log.Log().Error("序列化消息失败", "消息ID", msg.MessageID, "error", err)
			continue
		}

		msgKey := fmt.Sprintf(ConvMessagesKey, conversationID, msg.MessageID)
		pipe.Set(ctx, msgKey, string(msgJSON), time.Hour)

		// 2. 更新有序集合
		pipe.ZAdd(ctx, rankKey, &redis.Z{
			Score:  float64(msg.SendTime.Unix()),
			Member: fmt.Sprintf("%d", msg.MessageID),
		})
	}

	// 设置过期时间
	pipe.Expire(ctx, rankKey, time.Hour)

	// 执行缓存更新
	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Log().Error("更新消息缓存失败", "err", err)
	}
}
