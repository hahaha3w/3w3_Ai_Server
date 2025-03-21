package core

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

const (
	DefaultSQLLogFilePath = "sql.log"
)

var (
	db          *gorm.DB
	onceDb      sync.Once
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

// NewDB 初始化数据库连接
func NewDB(ctx context.Context) *gorm.DB {
	onceDb.Do(func() {
		var err error
		if db, err = gorm.Open(mysql.New(mysql.Config{
			DSN: fmt.Sprintf(viper.GetString("mysql.dsn"), viper.GetString("MYSQL_PASSWORD")),
		}), &gorm.Config{
			Logger:      getGormLogger(),
			PrepareStmt: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		}); err != nil {
			panic(err)
		}

	})
	return db.WithContext(ctx)
}

func getGormLogger() logger.Interface {
	ignoreRecordNotFound := false
	logLevel := logger.Info
	if !viper.GetBool("debug") {
		ignoreRecordNotFound = true
		logLevel = logger.Error
	}
	logFile, err := os.OpenFile(DefaultSQLLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	multiOutput := io.MultiWriter(os.Stdout, logFile)
	return logger.New(
		log.New(multiOutput, "[DB] ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: ignoreRecordNotFound,
			Colorful:                  false,
		},
	)
}
