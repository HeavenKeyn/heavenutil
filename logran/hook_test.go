package logran

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"testing"
)

func TestIOHook(t *testing.T) {
	config, err := LoadConfiguration("logran.xml")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for _, logger := range config.Logger {
		logrus.AddHook(NewIOHook(logger, os.Stdout))
	}
	logrus.SetOutput(ioutil.Discard)
	logrus.SetReportCaller(true)
	logrus.Info("info")
	logrus.Error("error")
}

func TestIOHook2(t *testing.T) {
	config, err := LoadConfiguration("logran.xml")
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
