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
	addAppender(config)
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
		for i, logger := range config.Logger {
			config.Logger[i].Func = strings.ReplaceAll(logger.Func, oldStr, newStr)
		}
	}
}

func addAppender(config *Configuration) {
	for _, appender := range config.Appender {
		for i, ref := range config.Root.AppenderRef {
			if ref.Ref == appender.Name {
				config.Root.AppenderRef[i].Appender = appender
			}
		}
		for i, logger := range config.Logger {
			for j, ref := range logger.AppenderRef {
				if ref.Ref == appender.Name {
					config.Logger[i].AppenderRef[j].Appender = appender
				}
			}
		}
	}
}
