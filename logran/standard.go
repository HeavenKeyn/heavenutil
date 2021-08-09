package logran

import (
	"github.com/HeavenKeyn/heavenutil/comutil"
	"github.com/HeavenKeyn/heavenutil/errutil"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func LoadStandardConfig(path string) error {
	logrus.SetLevel(logrus.AllLevels[len(logrus.AllLevels)-1])
	logrus.SetOutput(ioutil.Discard)
	logrus.SetReportCaller(true)
	config, err := LoadConfiguration(path)
	if err != nil {
		return err
	}

	for _, logger := range config.Logger {
		writer, err := getWriters(logger.AppenderRef)
		if err != nil {
			return err
		}
		logrus.AddHook(NewLoggerHook(logger, io.MultiWriter(writer)))
	}

	writer, err := getWriters(config.Root.AppenderRef)
	if err != nil {
		return err
	}
	logrus.AddHook(NewRootHook(*config, io.MultiWriter(writer)))
	return nil
}

func getWriters(refs []AppenderRef) (io.Writer, error) {
	writers := make([]io.Writer, 0)
	var errs errutil.MultiErrors
	for _, ref := range refs {
		writer, err := getWriter(ref)
		if err != nil {
			errs.AddError("创建日志文件失败", err)
		} else {
			writers = append(writers, writer)
		}
	}
	return io.MultiWriter(writers...), errs.GetError()
}

func getWriter(ref AppenderRef) (io.Writer, error) {
	if ref.Ref == "console" {
		return os.Stdout, nil
	} else {
		linkName, ok := ref.Appender.Policy["LinkName"]
		if !ok {
			linkName = ""
		}
		maxAge, err := comutil.ValueToInt64(ref.Appender.Policy["MaxAge"])
		if err != nil {
			maxAge = 24
		}
		rotationTime, err := comutil.ValueToInt64(ref.Appender.Policy["RotationTime"])
		if err != nil {
			rotationTime = 60
		}
		rl, err := rotatelogs.New(
			ref.Appender.File,
			rotatelogs.WithLinkName(linkName),
			rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Hour),
			rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Minute),
		)
		return rl, err
	}
}
