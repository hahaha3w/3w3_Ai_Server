package log

import "github.com/hahaha3w/3w3_Ai_Server/common/logs"

var logger *log.LogrusLogger

func RegisterLogger(path string, prefix string) {
	logger = log.NewLogrusLogger(path, prefix)
}
func Log() *log.LogrusLogger {
	return logger
}
