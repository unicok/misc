package log

import (
	"errors"

	"github.com/Sirupsen/logrus"
	"github.com/micro/go-os/log"
)

const (
	DefaultOutputName = "logrus"
)

type LogrusOption func(o *LogrusOptions)

type LogrusOptions struct {
	log.OutputOptions

	hooks     []logrus.Hook
	formatter logrus.Formatter
}

type output struct {
	*logrus.Logger

	opts LogrusOptions
	err  error
}

func (o *output) Send(e *log.Event) error {
	if o.Logger == nil {
		return o.err
	}

	fields := logrus.Fields{}
	for k, v := range e.Fields {
		fields[k] = v
	}

	entry := o.WithFields(fields)
	switch e.Level {
	case log.DebugLevel:
		entry.Debug(e.Message)
	case log.InfoLevel:
		entry.Info(e.Message)
	case log.WarnLevel:
		entry.Warn(e.Message)
	case log.ErrorLevel:
		entry.Error(e.Message)
	case log.FatalLevel:
		entry.Fatal(e.Message)
	}

	return nil
}

func (o *output) Flush() error {
	if o.Logger == nil {
		return o.err
	}
	return o.Flush()
}

func (o *output) Close() error {
	if o.Logger == nil {
		return o.err
	}
	return o.Close()
}

func (o *output) String() string {
	return "logrus"
}

func NewOutput(opts ...LogrusOption) log.Output {
	var options LogrusOptions
	for _, o := range opts {
		o(&options)
	}

	if len(options.Name) == 0 {
		options.Name = DefaultOutputName
	}

	l := logrus.New()

	//Log level is decided by the log of go-os
	l.Level = logrus.PanicLevel

	if options.formatter != nil {
		l.Formatter = options.formatter
	}

	for _, hook := range options.hooks {
		l.Hooks.Add(hook)
	}

	return &output{
		opts:   options,
		err:    errors.New("logrus initialize failed"),
		Logger: l,
	}
}

func WithFormatter(formatter logrus.Formatter) LogrusOption {
	return func(o *LogrusOptions) {
		o.formatter = formatter
	}
}

func WithHook(hook logrus.Hook) LogrusOption {
	return func(o *LogrusOptions) {
		o.hooks = append(o.hooks, hook)
	}
}
