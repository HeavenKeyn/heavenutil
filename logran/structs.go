package logran

import (
	"encoding/xml"
	"github.com/sirupsen/logrus"
	"io"
)

type Configuration struct {
	Property []Property `xml:"property"`
	Appender []Appender `xml:"appender"`
	Root     Root       `xml:"root"`
	Logger   []Logger   `xml:"logger"`
}

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

// Policy 该类主要为自定义日志规则
//如：日志文件名、文件大小、分割周期等
type Policy map[string]string

func (p *Policy) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	*p = Policy{}
	for {
		var e xmlMapEntry
		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		(*p)[e.XMLName.Local] = e.Value
	}
	return nil
}

type Appender struct {
	Name   string `xml:"name,attr"`
	File   string `xml:"file"`
	Policy Policy `xml:"policy"`
}

type Root struct {
	Level       logrus.Level  `xml:"level,attr"`
	AppenderRef []AppenderRef `xml:"appender-ref"`
}

type Logger struct {
	Func        string        `xml:"func,attr"`
	Level       logrus.Level  `xml:"level,attr"`
	Additivity  bool          `xml:"additivity,attr"`
	AppenderRef []AppenderRef `xml:"appender-ref"`
}

type AppenderRef struct {
	Ref      string `xml:"ref,attr"`
	Appender Appender
}
