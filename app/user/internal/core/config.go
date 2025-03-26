package core

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	DefaultConfigFileType = "toml"
	DefaultConfigFilePath = "conf/%s/conf.toml"
)

var ProviderSet = wire.NewSet(NewDB, NewContext, NewRedis)

func NewContext() context.Context {
	return context.Background()
}

type options struct {
	fileTypeOption string
	filePathOption string
}

type Option interface {
	apply(*options)
}

// o: immutable default config
var o = options{
	fileTypeOption: DefaultConfigFileType,
	filePathOption: DefaultConfigFilePath,
}

type fileTypeOption string

func (f fileTypeOption) apply(opt *options) {
	opt.fileTypeOption = string(f)
}
func WithFileTypeOption(fileType string) Option {
	return fileTypeOption(fileType)
}

type filePathOption string

func (f filePathOption) apply(opt *options) {
	opt.filePathOption = string(f)
}
func WithFilePathOption(fileType string) Option {
	return filePathOption(fileType)
}

func LoadConfig(opts ...Option) (err error) {
	var mode string
	for _, opt := range opts {
		opt.apply(&o)
	}
	err = godotenv.Load()
	if err != nil {
		log.Println("no .env file found,use default")
	}
	viper.AutomaticEnv()
	if mode = viper.GetString("MODE"); mode == "" {
		mode = "test"
	}
	confPath := fmt.Sprintf(o.filePathOption, mode)
	if _, err := os.Stat(confPath); err != nil {
		return fmt.Errorf(
			"load config: %w", err)
	}

	viper.SetConfigFile(confPath)
	viper.SetConfigType(o.fileTypeOption)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf(
			"read config: %w", err)
	}
	return nil
}
