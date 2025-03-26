package core

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	redisClient *redis.Client
	onceRedis   sync.Once
)

// NewRedis 初始化 Redis 客户端
func NewRedis(ctx context.Context) *redis.Client {
	onceRedis.Do(func() {
		// 从配置文件中读取 Redis 连接信息
		addr := viper.GetString("redis.addr")
		password := viper.GetString("redis.password")
		db := viper.GetInt("redis.db")

		// 初始化 Redis 客户端
		redisClient = redis.NewClient(&redis.Options{
			Addr:     addr,     // Redis 地址
			Password: password, // Redis 密码
			DB:       db,       // Redis 数据库编号
		})

		// 测试连接
		if err := redisClient.Ping(ctx).Err(); err != nil {
			panic(fmt.Sprintf("failed to connect to Redis: %v", err))
		}

		log.Println("Redis connected successfully!")
	})
	return redisClient
}
