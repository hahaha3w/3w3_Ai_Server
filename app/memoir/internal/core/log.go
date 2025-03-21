package core

import "github.com/hahaha3w/3w3_Ai_Server/memoir/pkg/log"

const (
	logPath   = "logs"
	logPrefix = "memoir"
)

func LoadLogger() {
	log.RegisterLogger(logPath, logPrefix)
}
