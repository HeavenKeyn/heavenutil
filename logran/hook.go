package logran

import (
	"github.com/sirupsen/logrus"
	"io"
)

type IOHook struct {
	Logger
	levels []logrus.Level
	Writer io.Writer
}

func NewIOHook(logger Logger, writer io.Writer) IOHook {
	ioHook := IOHook{
		Logger: logger,
		Writer: writer,
	}
	levels := make([]logrus.Level, 0)
	for _, level := range logrus.AllLevels {
		if level <= logger.Level {
			levels = append(levels, level)
		}
	}
	ioHook.levels = levels
	return ioHook
}

func (hook IOHook) Levels() []logrus.Level {
	return hook.levels
}

func (hook IOHook) Fire(entry *logrus.Entry) error {
	if entry.Caller.Function != hook.Func {
		return nil
	}
	line, err := entry.Bytes()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write(line)
	return err
}
