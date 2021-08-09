package logran

import (
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestIOHook(t *testing.T) {
	config, err := LoadConfiguration("testdata/logran.xml")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for _, logger := range config.Logger {
		writers := make([]io.Writer, 0)
		for _, ref := range logger.AppenderRef {
			if ref.Ref == "console" {
				writers = append(writers, os.Stdout)
			} else {
				fp, err := os.OpenFile(ref.Appender.File, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
				if err != nil {
					t.Error("创建日志文件失败", err)
				} else {
					writers = append(writers, fp)
				}
			}
		}
		logrus.AddHook(NewIOHook(logger, io.MultiWriter(writers...)))
	}
	logrus.SetOutput(ioutil.Discard)
	logrus.SetReportCaller(true)
	logrus.Info("info")
	logrus.Error("error")
}

func TestIOHook2(t *testing.T) {
	config, err := LoadConfiguration("testdata/logran.xml")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for _, logger := range config.Logger {
		logrus.AddHook(IOHook{
			Logger: logger,
			Writer: os.Stderr,
		})
	}
	logrus.SetOutput(ioutil.Discard)
	logrus.SetReportCaller(true)
	logrus.Info("aaaaa")
}
