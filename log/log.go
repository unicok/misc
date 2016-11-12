package log

import (
	"fmt"

	"github.com/micro/cli"
	"github.com/micro/go-micro/cmd"
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
	app := cmd.App()

	app.Flags = append(app.Flags,
		cli.StringFlag{
			Name:   "log_level",
			EnvVar: "UNI_LOG_LEVEL",
			Usage:  "The level of logging. Value:[debug, info, warn, error, fatal]",
			Value:  "info",
		},
	)

	before := app.Before

	app.Before = func(ctx *cli.Context) error {
		if lv := ctx.String("log_level"); lv != "" {
			switch lv {
			case "debug":
				level = log.DebugLevel
			case "info":
				level = log.InfoLevel
			case "warn":
				level = log.WarnLevel
			case "error":
				level = log.ErrorLevel
			case "fatal":
				level = log.FatalLevel
			}
		}

		logger = log.NewLog(
			log.WithLevel(level),
			log.WithOutput(NewOutput()),
		)

		logger.Info("Log Level:", level)

		return before(ctx)
	}
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
