package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

//write the log to terminal
type LogLevel uint16

const (
	UNKNOW LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	switch strings.ToLower(s) {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "error":
		return ERROR, nil
	case "trace":
		return TRACE, nil
	case "warning":
		return WARNING, nil
	case "fatal":
		return FATAL, nil
	default:
		return UNKNOW, errors.New("not support log level")
	}
}

type ConsoleLogger struct {
	Level LogLevel
}

func NewLog(s string) ConsoleLogger {
	level, err := parseLogLevel(s)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (l ConsoleLogger) Debug(format string, a ...interface{}) {
	if l.Level <= DEBUG {
		log(ERROR, format, a...)
	}

}

func (l ConsoleLogger) Info(format string, a ...interface{}) {
	if l.Level <= INFO {
		log(ERROR, format, a...)
	}
}

func (l ConsoleLogger) Error(format string, a ...interface{}) {
	if l.Level <= ERROR {
		log(ERROR, format, a...)
	}
}

func log(level LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	fileName, funcName, lineNo := sourceInfo(3)
	_, logLevel, _ := sourceInfo(2)
	fmt.Printf("[%s] [%s] [%s|%s|%d] %s\n", now.Format("2006-01-02 15:04:05"), strings.ToUpper(logLevel), fileName, funcName, lineNo, msg)
}

func sourceInfo(skip int) (fileName, funcName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("call runtime.Caller failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	splitNames := strings.Split(funcName, ".")
	funcName = splitNames[len(splitNames)-1]
	return
}
