package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/internal/infra/rpc"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/pkg/count"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/pkg/llm"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/pkg/log"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/activity"
	"time"
)

var _ domain.MemoirUsecase = &ConcreteMemoirUsecase{}

const (
	CacheKeyMemoirListPrefix   = "memoir:list:%d"      // 用户回忆录列表缓存key，参数为用户ID
	CacheKeyMemoirDetailPrefix = "memoir:detail:%d"    // 回忆录详情缓存key，参数为回忆录ID
	CacheExpiration            = 12 * time.Hour        // 缓存过期时间
	DateLayout                 = "2006-01-02 15:04:05" // 时间格式
	EmptyDataExpiration        = 1 * time.Minute       // 空数据缓存时间
)

// CachedList 缓存数据结构体
type CachedList struct {
	Data  []*domain.Memoir `json:"data"`
	Total int32            `json:"total"`
}

type ConcreteMemoirUsecase struct {
	repo  domain.MemoirRepo
	cache *redis.Client
}

func NewConcreteMemoirUsecase(repo domain.MemoirRepo, cache *redis.Client) *ConcreteMemoirUsecase {
	return &ConcreteMemoirUsecase{
		repo:  repo,
		cache: cache,
	}
}

// GenerateMemoir 生成回忆录
func (c ConcreteMemoirUsecase) GenerateMemoir(ctx context.Context, userID int, style, memoirType, startDate, endDate string) (*domain.Memoir, error) {
	// 将字符串类型转化成 time.Time 类型
	var startTime, endTime time.Time
	now := time.Now()

	const DateLayout = "2006-01-02"

	//locEast8, err := time.LoadLocation("Asia/Shanghai")
	//if err != nil {
	//	return nil, fmt.Errorf("failed to load East 8 location: %w", err)
	//}
	var err error

	if startDate != "" && endDate != "" {
		startTime, err = time.Parse(time.DateOnly, startDate)
		parsedStart, err := time.Parse(DateLayout, startDate)
		if err != nil {
			log.Log().Error(err)
			return nil, errors.New("invalid startDate format, expected YYYY-MM-DD")
		}
		startTime = parsedStart

		endTime, err = time.Parse(time.DateOnly, endDate)
		parsedEnd, err := time.Parse(DateLayout, endDate)
		if err != nil {
			log.Log().Error(err)
			return nil, errors.New("invalid endDate format, expected YYYY-MM-DD")
		}
		endTime = parsedEnd
	}
	questions, err := c.repo.GetUserMessages(ctx, userID, startTime, endTime)
	if err != nil {
		log.Log().Error(err)
		return nil, err
	}
	if len(questions) == 0 {
		log.Log().Error("no question found")
		return nil, errors.New("no question found")
	}
	title, content, err := llm.GenerateMemoir(ctx, questions, style)
	if err != nil {
		log.Log().Error(err)
		return nil, err
	}
	// 创建回忆录对象
	memoir := &domain.Memoir{
		UserID:    userID,
		Title:     title,
		Content:   content,
		Type:      memoirType,
		Style:     style,
		StartDate: startTime,
		EndDate:   endTime,
		CreatedAt: now,
	}

	// 调用 repo 层的 CreateMemoir 方法保存回忆录
	if err := c.repo.CreateMemoir(ctx, memoir); err != nil {
		return nil, err
	}

	// 清理相关缓存
	patterns := []string{
		fmt.Sprintf(CacheKeyMemoirListPrefix+":*", userID),   // 列表缓存
		fmt.Sprintf(CacheKeyMemoirDetailPrefix+":*", userID), // 详情缓存
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

	// 修改用户次数统计
	userCount := &domain.UpdateUserCount{
		UserId:          int32(userID),
		ChatCountIncr:   0,
		MemoirCountIncr: 1,
		UseDaysIncr:     0,
	}
	_, err = count.UpdateUserCount(userCount, c.cache)
	if err != nil {
		log.Log().Error("Failed to update user count: %v", err)
		return nil, fmt.Errorf("internal error: %w", err)
	}

	// 调用 rpc
	_, err = rpc.ActivityClient.CreateUserActivity(ctx, &activity.CreateUserActivityReq{
		UserId:      int64(userID),
		RelationId:  int64(memoir.MemoirID),
		Type:        "memoir",
		Description: "AI为你创建了\"" + memoir.Title + "\"日记",
	})
	if err != nil {
		log.Log().Error("Failed to create activity: %v", err)
		return nil, err
	}

	return memoir, nil
}

// GetMemoirList 获取回忆录列表
func (c ConcreteMemoirUsecase) GetMemoirList(ctx context.Context, userID int, memoirType, startDate, endDate string, page, pageSize int32) ([]*domain.Memoir, bool, error) {
	//// 规范化缓存键参数
	//
	////TODO 处理style
	//params := struct {
	//	Type      string
	//	Style     string
	//	StartDate string
	//	EndDate   string
	//	Page      int32
	//	PageSize  int32
	//}{
	//	Type:      memoirType,
	//	Style:     "",
	//	StartDate: normalizeDate(startDate),
	//	EndDate:   normalizeDate(endDate),
	//	Page:      page,
	//	PageSize:  pageSize,
	//}
	//
	//cacheKey := fmt.Sprintf("%s:%v",
	//	fmt.Sprintf(CacheKeyMemoirListPrefix, userID),
	//	md5.Sum([]byte(fmt.Sprintf("%+v", params))), // 使用结构体哈希作为唯一标识
	//)
	//
	//// 尝试获取完整缓存
	//var cached CachedList
	//cachedData, err := c.cache.Get(ctx, cacheKey).Bytes()
	//if err == nil {
	//	if err := json.Unmarshal(cachedData, &cached); err == nil {
	//		if len(cached.Data) > 0 || cached.Total > 0 {
	//			return cached.Data, cached.Total, nil
	//		}
	//		return []*domain.Memoir{}, 0, nil
	//	}
	//}

	// 数据库查询
	memoirs, _, err := c.repo.GetMemoirsByUserID(ctx, userID, "", memoirType, startDate, endDate, page, pageSize)
	if err != nil {
		return nil, false, err
	}
	var hasMore bool
	if len(memoirs) > int(pageSize) {
		hasMore = true
	}

	//// 构建缓存对象
	//cacheObj := CachedList{
	//	Data:  memoirs,
	//	Total: total,
	//}

	//expiration := CacheExpiration
	//if len(memoirs) == 0 && total == 0 {
	//	expiration = EmptyDataExpiration
	//}
	//
	//// 统一存储缓存
	//cacheData, err := json.Marshal(cacheObj)
	//if err != nil {
	//	log.Log().Error("Failed to marshal cache data: %v", err)
	//} else {
	//	if err := c.cache.Set(ctx, cacheKey, cacheData, expiration).Err(); err != nil {
	//		log.Log().Error("Failed to cache memoir list: %v", err)
	//	}
	//}

	if len(memoirs) == 0 {
		return memoirs, false, nil
	}
	return memoirs[:len(memoirs)-1], hasMore, nil
}

// GetMemoirDetail 获取回忆录详情
func (c ConcreteMemoirUsecase) GetMemoirDetail(ctx context.Context, memoirID, userID int) (*domain.Memoir, error) {
	cacheKey := fmt.Sprintf(CacheKeyMemoirDetailPrefix, memoirID)

	// 从缓存中获取回忆录详情
	var memoir domain.Memoir
	cachedData, err := c.cache.Get(ctx, cacheKey).Bytes()
	if err == nil {
		if err := json.Unmarshal(cachedData, &memoir); err == nil {
			return &memoir, nil
		}
	}

	// 如果缓存中没有，则从数据库中获取
	memoirPtr, err := c.repo.GetMemoirByID(ctx, memoirID, userID)
	if err != nil {
		return nil, err
	}

	// 将回忆录详情缓存起来
	cacheData, err := json.Marshal(memoirPtr)
	if err != nil {
		log.Log().Error("Failed to marshal cache data: %v", err)
	} else {
		if err := c.cache.Set(ctx, cacheKey, cacheData, CacheExpiration).Err(); err != nil {
			log.Log().Error("Failed to cache memoir detail: %v", err)
		}
	}

	return memoirPtr, nil
}

// DeleteMemoir 删除回忆录
func (c ConcreteMemoirUsecase) DeleteMemoir(ctx context.Context, memoirID, userID int) error {
	// 删除回忆录
	affected, err := c.repo.DeleteMemoir(ctx, memoirID, userID)
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("memoir not found")
	}

	// 删除回忆录详情缓存
	cacheKey := fmt.Sprintf(CacheKeyMemoirDetailPrefix, memoirID)
	if err := c.cache.Del(ctx, cacheKey).Err(); err != nil {
		log.Log().Error("Failed to delete memoir detail cache: %v", err)
	}

	// 删除回忆录列表缓存
	pattern := fmt.Sprintf(CacheKeyMemoirListPrefix+":*", userID)
	keys, err := c.cache.Keys(ctx, pattern).Result()
	if err == nil && len(keys) > 0 {
		if err := c.cache.Del(ctx, keys...).Err(); err != nil {
			log.Log().Error("Failed to delete memoir list caches: %v", err)
		}
	}

	// 修改用户次数统计
	userCount := &domain.UpdateUserCount{
		UserId:          int32(userID),
		ChatCountIncr:   0,
		MemoirCountIncr: -1,
		UseDaysIncr:     0,
	}
	_, err = count.UpdateUserCount(userCount, c.cache)
	if err != nil {
		log.Log().Error("Failed to update user count: %v", err)
		return err
	}

	return nil
}

// normalizeDate 规范化日期参数
func normalizeDate(date string) string {
	if date == "" {
		return "nil"
	}
	return date
}
