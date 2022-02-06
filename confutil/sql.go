package confutil

import (
	"fmt"
	"strings"
)

type DBType string

const (
	MySQL      DBType = "mysql"
	Clickhouse DBType = "clickhouse"
)

type SQLdb struct {
	Type     DBType
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Network  string
	DSN      string
}

func (p *SQLdb) GetDSN() string {
	if p.DSN != "" { //如果已配置DSN，则直接返回
		return p.DSN
	}
	if p.Network == "" { //如果未配置network，则默认tcp
		p.Network = "tcp"
	}
	switch p.Type { //根据不同的数据库，生成dsn
	case MySQL:
		p.DSN = fmt.Sprintf("%s:%s@%s(%s:%d)/%s", p.Username, p.Password, p.Network, p.Host, p.Port, p.Database) +
			"?charset=utf8mb4&parseTime=True&loc=Local"
	case Clickhouse:
		var dsn strings.Builder
		dsn.WriteString(fmt.Sprintf("%s://%s:%d?database=%s", p.Network, p.Host, p.Port, p.Database))
		if p.Username != "" { //Clickhouse支持默认用户及空密码
			dsn.WriteString(fmt.Sprintf("&username=%s&password=%s", p.Username, p.Password))
		}
		dsn.WriteString("&charset=utf8&loc=Local&parseTime=true")
		p.DSN = dsn.String()
	default:

	}
	return p.DSN
}
