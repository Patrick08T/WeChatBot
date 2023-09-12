package log

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var logger = log.New()
var logLevel = map[string]log.Level{
	"PANIC": log.PanicLevel,
	"FATAL": log.FatalLevel,
	"ERROR": log.ErrorLevel,
	"WARN":  log.WarnLevel,
	"INFO":  log.InfoLevel,
	"DEBUG": log.DebugLevel,
	"TRACE": log.TraceLevel,
}

func InitLogger(level string, formatter log.Formatter, filename string) {
	logger.SetLevel(logLevel[level])
	logger.SetFormatter(formatter)
	logger.SetOutput(io.MultiWriter(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	}, os.Stdout))
}

func PANIC(v ...interface{}) {
	logger.Panic(v)
}

func FATAL(v ...interface{}) {
	logger.Fatal(v)
}

func ERROR(v ...interface{}) {
	logger.Error(v)
}

func WARN(v ...interface{}) {
	logger.Warn(v)
}

func INFO(v ...interface{}) {
	logger.Info(v)
}

func DEBUG(v ...interface{}) {
	logger.Debug(v)
}

func PANICF(format string, v ...interface{}) {
	logger.Panicf(format, v...)
}
func FATALF(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}
func ERRORF(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}
func WARNF(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}
func INFOF(format string, v ...interface{}) {
	logger.Infof(format, v...)
}
func DEBUGF(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}
