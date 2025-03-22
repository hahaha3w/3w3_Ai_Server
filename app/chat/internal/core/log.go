package core

import (
	"github.com/hahaha3w/3w3_Ai_Server/chat/pkg/log"
	logger "github.com/hahaha3w/3w3_Ai_Server/common/logs"
)

const (
	logPath   = "logs"
	logPrefix = "user"
)

func LoadLogger() {
	log.RegisterLogger(logger.NewLogrusLogger(logPath, logPrefix))
}
