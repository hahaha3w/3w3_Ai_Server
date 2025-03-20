package core

import (
	"github.com/hahaha3w/3w3_Ai_Server/user/pkg/log"
)

const (
	logPath   = "logs"
	logPrefix = "user"
)

func LoadLogger() {
	log.RegisterLogger(logger.NewLogrusLogger(logPath, logPrefix))
}
