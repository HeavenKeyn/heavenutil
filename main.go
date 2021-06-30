package main

func main() {
	//var mdb mysql.MDB
	//mdb.Connect("heaven.com","root","20210618")
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
