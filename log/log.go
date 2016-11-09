package log

import (
	"fmt"
	"os"
	"strconv"

	"github.com/micro/go-os/log"
)

const (
	levelEnv     = "LOG_LEVEL"
	defaultLevel = log.InfoLevel
)

var (
	level  log.Level
	logger log.Log
)

func init() {
	level = defaultLevel
	if env := os.Getenv(levelEnv); env != "" {
		v, err := strconv.ParseUint(env, 10, 32)
		if err == nil {
			level = log.Level(v)
		}
	}

	logger = log.NewLog(
		log.WithLevel(log.InfoLevel),
		// log.WithFields(log.Fields{
		// 	"logger": "platform",
		// }),
		log.WithOutput(
			NewOutput(),
		),
	)

	logger.Info("Log Level:", level)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Panic(args ...interface{}) {
	logger.Error(args...)
	panic(fmt.Sprint(args...))
}

// Formatted logger
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	Panic(fmt.Sprintf(format, args...))
}

// Specify your own levels
func Log(l log.Level, args ...interface{}) {
	logger.Log(l, args...)
}

func Logf(l log.Level, format string, args ...interface{}) {
	logger.Logf(l, format, args...)
}

// Returns with extra fields
func WithFields(f log.Fields) log.Logger {
	return logger.WithFields(f)
}
