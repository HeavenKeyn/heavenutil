package logran

import (
	"github.com/HeavenKeyn/heavenutil/errutil"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
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
		if ref.Ref == "console" {
			writers = append(writers, os.Stdout)
		} else {
			fp, err := os.OpenFile(ref.Appender.File, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
			if err != nil {
				errs.AddError("创建日志文件失败", err)
			} else {
				writers = append(writers, fp)
			}
		}
	}
	return io.MultiWriter(writers...), errs.GetError()
}
