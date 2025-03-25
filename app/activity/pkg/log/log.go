package log

import logger "github.com/hahaha3w/3w3_Ai_Server/common/logs"

var log *logger.LogrusLogger

func RegisterLogger(logger *logger.LogrusLogger) {
	log = logger
}

func Log() *logger.LogrusLogger {
	if log == nil {
		panic("implement not found for interface Logger, please register")
	}
	return log
}
