package properties

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadProperties(path string, out interface{}) error {
	doc, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(doc, out)
	if err != nil {
		return err
	}
	return nil
}
