package logger

import (
	"fmt"
	"github.com/liu-cn/runbox/pkg/jsonx"
	"runtime"
	"strconv"
	"time"
)

type Logger struct {
	Level   string
	DataMap map[string]interface{} `json:"data_map"`
}

func (l *Logger) logPrint(level string, msg string) {
	_, file, line, _ := runtime.Caller(2)
	l.DataMap["a_ts"] = time.Now().Unix()
	l.DataMap["a_msg"] = msg
	l.DataMap["a_stack"] = file + ":" + strconv.Itoa(line)
	l.DataMap["a_level"] = level

	fmt.Println("<Logger>" + jsonx.String(l.DataMap) + "</Logger>")
}

func (l *Logger) logPrintf(level string, formatMsg string, args any) {
	_, file, line, _ := runtime.Caller(1)
	l.DataMap["a_ts"] = time.Now().Unix()
	l.DataMap["a_msg"] = fmt.Sprintf(formatMsg, args)
	l.DataMap["a_stack"] = file + ":" + strconv.Itoa(line)
	if l.Level != "" {
		l.DataMap["a_level"] = l.Level
	} else {
		l.DataMap["a_level"] = level
	}
	fmt.Println("<Logger>" + jsonx.String(l.DataMap) + "</Logger>")
}

// SetLevel 可以自定义level
func (l *Logger) SetLevel(level string) *Logger {
	l.Level = level
	return l
}
func (l *Logger) Info(msg string) {
	l.logPrint("INFO", msg)
}
func (l *Logger) Infof(formatMsg string, args ...any) {
	if len(args) > 0 {
		l.logPrintf("INFO", formatMsg, args)
	} else {
		l.logPrint("INFO", formatMsg)
	}
}

func (l *Logger) Warn(msg string) {
	l.logPrint("WARN", msg)
}
func (l *Logger) Warnf(formatMsg string, args ...any) {
	if len(args) > 0 {
		l.logPrintf("WARN", formatMsg, args)
	} else {
		l.logPrint("WARN", formatMsg)
	}
}

func (l *Logger) Error(msg string) {
	l.logPrint("ERROR", msg)
}
func (l *Logger) Errorf(formatMsg string, args ...any) {
	if len(args) > 0 {
		l.logPrintf("ERROR", formatMsg, args)
	} else {
		l.logPrint("ERROR", formatMsg)
	}
}
func (l *Logger) Panic(msg string) {
	l.logPrint("PANIC", msg)
}
func (l *Logger) Panicf(formatMsg string, args ...any) {
	if len(args) > 0 {
		l.logPrintf("PANIC", formatMsg, args)
	} else {
		l.logPrint("PANIC", formatMsg)
	}
}
func (l *Logger) Debug(msg string) {
	l.logPrint("DEBUG", msg)
}
func (l *Logger) Debugf(formatMsg string, args ...any) {
	if len(args) > 0 {
		l.logPrintf("DEBUG", formatMsg, args)
	} else {
		l.logPrint("DEBUG", formatMsg)
	}
}

func (l *Logger) Fatal(msg string) {
	l.logPrint("FATAL", msg)
}

func (l *Logger) Fatalf(formatMsg string, args ...any) {
	if len(args) > 0 {
		l.logPrintf("FATAL", formatMsg, args)
	} else {
		l.logPrint("FATAL", formatMsg)
	}
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	if _, ok := l.DataMap[key]; !ok {
		l.DataMap[key] = value
	} else {
		l.DataMap["runner."+key] = value
	}
	return l
}
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	for k, v := range fields {
		if _, ok := l.DataMap[k]; !ok {
			l.DataMap[k] = v
		} else {
			l.DataMap["runner."+k] = v
		}
	}
	return l
}
