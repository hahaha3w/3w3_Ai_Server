package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogrusLogger struct {
	*logrus.Logger
}
type formatter struct {
	prefix string
}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", t.prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n", t.prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

type hook struct {
	logPath string
	prefix  string
	file    *os.File
}

func (h *hook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (h *hook) Fire(entry *logrus.Entry) (err error) {
	timer := entry.Time.Format("2006-01-02_15-04")
	line, _ := entry.String()
	// 时间不等
	defer func() { _ = h.file.Close() }()
	err = os.MkdirAll(fmt.Sprintf("%s/%s", h.logPath, timer), os.ModePerm)
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("%s/%s/%s.log", h.logPath, timer, h.prefix)

	h.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	_, err = h.file.Write([]byte(line))
	return err
}
func NewLogrusLogger(logPath string, logPrefix string) *LogrusLogger {
	var level = logrus.InfoLevel
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetReportCaller(true)
	l.SetFormatter(&formatter{
		prefix: logPrefix,
	})
	l.AddHook(
		&hook{
			logPath: logPath,
			prefix:  logPrefix,
		})
	if mode := os.Getenv("mode"); mode == "" {
		level = logrus.DebugLevel
	}
	l.SetLevel(level)
	return &LogrusLogger{l}
}
