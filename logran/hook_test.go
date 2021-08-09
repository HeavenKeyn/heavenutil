package logran

import (
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestIOHook(t *testing.T) {
	logrus.SetLevel(logrus.AllLevels[len(logrus.AllLevels)-1])
	logrus.SetOutput(ioutil.Discard)
	logrus.SetReportCaller(true)
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
		logrus.AddHook(NewLoggerHook(logger, io.MultiWriter(writers...)))
	}

	writers := make([]io.Writer, 0)
	for _, ref := range config.Root.AppenderRef {
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
	logrus.AddHook(NewRootHook(*config, io.MultiWriter(writers...)))

	logrus.Debug("debug")
	logrus.Info("info")
	logrus.Error("error")
	testIOHook1()
	testIOHook2()
}

func testIOHook1() {
	logrus.Debug("debug")
	logrus.Info("info")
	logrus.Error("error")
}

func testIOHook2() {
	logrus.Debug("debug")
	logrus.Info("info")
	logrus.Error("error")
}
