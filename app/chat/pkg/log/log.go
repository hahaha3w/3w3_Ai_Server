package log

import logger "github.com/hahaha3w/3w3_Ai_Server/common/logs"

var log logger.Logger

func RegisterLogger(logger logger.Logger) {
	log = logger
}
func Log() logger.Logger {
	if log == nil {
		panic("implement not found for interface Logger, please register")
	}
	return log
}
