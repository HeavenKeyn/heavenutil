package main

import (
	"fmt"
	"github.com/HeavenKeyn/heavenutil/comutil"
)

func main() {
	fmt.Println(comutil.HumpToUnderline("MonMoo"))
}

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
