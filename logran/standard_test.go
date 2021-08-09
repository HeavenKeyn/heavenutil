package logran

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLoadStandardConfig(t *testing.T) {
	err := LoadStandardConfig("testdata/logran.xml")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	logrus.Debug("debug")
	logrus.Info("info")
	logrus.Error("error")
	testIOHook1()
	testIOHook2()
}
