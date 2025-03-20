package core

import (
	"context"
	"fmt"
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
	db     *gorm.DB
	onceDb sync.Once
)

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
