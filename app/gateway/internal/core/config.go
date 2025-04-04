package core

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

const (
	DefaultConfigFileType = "toml"
	DefaultConfigFilePath = "conf/%s/conf.toml"
)

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
