package log

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type myFormatter struct {}
var levelList = [] string{"PANIC", "FATAL", "ERROR", "WARN", "INFO", "DEBUG", "TRACE"}
var modes = map[string]int{"PANIC": 0, "FATAL": 1, "ERROR": 2, "WARN": 3, "INFO": 4, "DEBUG": 5, "TRACE": 6}

func (mf *myFormatter) Format(entry *logrus.Entry) ([]byte, error){
    var b *bytes.Buffer
    if entry.Buffer != nil {
        b = entry.Buffer
    } else {
        b = &bytes.Buffer{}
    }
    level := levelList[int(entry.Level)]
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = 32 // green
	case logrus.WarnLevel:
		levelColor = 33 // yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
    strList := strings.Split(entry.Caller.File, "/")
    fileName := strList[len(strList)-1]
    b.WriteString(fmt.Sprintf("%s  [%s:%-3d] \x1b[%dm%-6s\x1b[0m  %s\n",
        entry.Time.Format("2006-01-02 15:04:05"), fileName,
        entry.Caller.Line, levelColor, level, entry.Message))
    return b.Bytes(), nil
}

// New Creation a new logger
func New(level string) *logrus.Logger{
	if level != "" {
		level = strings.ToUpper(level)
	} else {
		level = "trace"
	}
	mode := modes[level]
	logger := logrus.New()
	logger.SetLevel(logrus.Level(mode))
	logger.SetReportCaller(true)
    logger.SetFormatter(&myFormatter{})
	return logger
}