package log

var logger *LogrusLogger

func RegisterLogger(path string, prefix string) {
	logger = NewLogrusLogger(path, prefix)
}
func Log() *LogrusLogger {
	return logger
}
