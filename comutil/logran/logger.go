package logran

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

func LoadConfiguration(path string) (*Configuration, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := new(Configuration)
	err = xml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}
	replaceProperty(config)
	return config, nil
}

func replaceProperty(config *Configuration) {
	for _, property := range config.Property {
		oldStr := fmt.Sprintf("${%s}", property.Name)
		newStr := property.Value
		for i, appender := range config.Appender {
			config.Appender[i].File = strings.ReplaceAll(appender.File, oldStr, newStr)
			for k, v := range appender.Policy {
				config.Appender[i].Policy[k] = strings.ReplaceAll(v, oldStr, newStr)
			}
		}
	}

}
