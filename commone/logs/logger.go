package log

type Logger interface {
	Debug(msg any, fields ...Field)
	Info(msg any, fields ...Field)
	Warn(msg any, fields ...Field)
	Error(msg any, fields ...Field)
	Fatal(msg any, fields ...Field)
	Panic(msg any, fields ...Field)
}

type Field struct {
	Key   string
	Value interface{}
}
