package logran

import "testing"

func TestLoadConfiguration(t *testing.T) {
	config, err := LoadConfiguration("logran.xml")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(config)
	}
}
