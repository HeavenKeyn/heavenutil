package confutil

import (
	"flag"
	"os"
	"strings"
	"time"
)

type ConfigClient struct {
}

var DefaultClient ConfigClient

const (
	Split = ","
)

func (c ConfigClient) String(name string, value string, usage string) string {
	val, ok := os.LookupEnv(getENVName(name))
	if ok {
		return val
	}
	flagVal := flag.String(getFlagName(name), value, usage)
	if flagVal != nil {
		return *flagVal
	}
	return value
}

func String(name string, value string, usage string) string {
	return DefaultClient.String(name, value, usage)
}

func (c ConfigClient) Array(name string, value []string, usage string) []string {
	s := String(name, strings.Join(value, Split), usage)
	return strings.Split(s, Split)
}

func Array(name string, value []string, usage string) []string {
	return DefaultClient.Array(name, value, usage)
}

func (c ConfigClient) Duration(name string, value time.Duration, usage string) time.Duration {
	val, ok := os.LookupEnv(getENVName(name))
	if ok {
		d, err := time.ParseDuration(val)
		if err != nil {
			return value
		}
		return d
	}
	flagVal := flag.Duration(getFlagName(name), value, usage)
	if flagVal != nil {
		return *flagVal
	}
	return value
}

func Duration(name string, value time.Duration, usage string) time.Duration {
	return DefaultClient.Duration(name, value, usage)
}

func getENVName(name string) string {
	return strings.ToUpper(name)
}

func getFlagName(name string) string {
	return strings.ToLower(name)
}

func (c ConfigClient) Parse() {
	flag.Parse()
}

func Parse() {
	DefaultClient.Parse()
}
