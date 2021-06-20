package properties

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type Properties struct {
	InfluxDB InfluxDB
	MySQL    MySQL
}

type InfluxDB struct {
	Addr     string
	Database string
	Username string
	Password string
}

type MySQL struct {
	Username string
	Password string
	Network  string
	Host     string
	Port     int
	Database string
}

func LoadProperties(path string, out *interface{}) error {
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
