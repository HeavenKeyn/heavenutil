package logran

import (
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	config, err := LoadConfiguration("testdata/logran.xml")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(config)
	}
}
