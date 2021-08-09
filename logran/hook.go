package logran

import (
	"github.com/sirupsen/logrus"
	"io"
)

type IOHook struct {
	levels []logrus.Level
	Writer io.Writer
}

func (h *IOHook) SetLevel(level logrus.Level) {
	levels := make([]logrus.Level, 0)
	for _, l := range logrus.AllLevels {
		if l <= level {
			levels = append(levels, l)
		}
	}
	h.levels = levels
}

func (h *IOHook) SetWriter(writers ...io.Writer) {
	h.Writer = io.MultiWriter(writers...)
}

func (h IOHook) Levels() []logrus.Level {
	return h.levels
}

func (h IOHook) Fire(entry *logrus.Entry) error {
	line, err := entry.Bytes()
	if err != nil {
		return err
	}
	_, err = h.Writer.Write(line)
	return err
}

type RootHook struct {
	IOHook
	funcS []string
}

func NewRootHook(config Configuration, writer io.Writer) RootHook {
	var hook RootHook
	hook.SetWriter(writer)
	hook.SetLevel(config.Root.Level)
	hook.funcS = make([]string, 0)
	for _, logger := range config.Logger {
		if !logger.Additivity {
			hook.funcS = append(hook.funcS, logger.Func)
		}
	}
	return hook
}

func (hook RootHook) Fire(entry *logrus.Entry) error {
	for _, s := range hook.funcS {
		if entry.Caller.Function == s {
			return nil
		}
	}
	return hook.IOHook.Fire(entry)
}

type LoggerHook struct {
	IOHook
	logger Logger
}

func NewLoggerHook(logger Logger, writer io.Writer) LoggerHook {
	ioHook := LoggerHook{
		logger: logger,
	}
	ioHook.SetLevel(logger.Level)
	ioHook.SetWriter(writer)
	return ioHook
}

func (hook LoggerHook) Fire(entry *logrus.Entry) error {
	if entry.Caller.Function != hook.logger.Func {
		return nil
	}
	return hook.IOHook.Fire(entry)
}
