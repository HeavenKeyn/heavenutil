package confutil

import (
	"fmt"
	"strings"
)

type DBType string

const (
	MySQL      DBType = "mysql"
	Clickhouse DBType = "clickhouse"
	MongoDB    DBType = "mongodb"
)

// DBProp 数据库配置类，根据配置生成DSN或直接配置DSN
type DBProp struct {
	Type     DBType `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Network  string `yaml:"network"`
	DSN      string `yaml:"dsn"`
}

func (p DBProp) GetType() DBType {
	return p.Type
}

func (p DBProp) GetHost() string {
	return p.Host
}

func (p DBProp) GetPort() int {
	return p.Port
}

func (p DBProp) GetUsername() string {
	return p.Username
}

func (p DBProp) GetPassword() string {
	return p.Password
}

func (p DBProp) GetDatabase() string {
	return p.Database
}

func (p DBProp) GetNetwork() string {
	return p.Network
}

func (p *DBProp) GetDSN() string {
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
	case MongoDB:
		p.DSN = fmt.Sprintf("%s://%s:%s@%s:%d", MongoDB, p.Username, p.Password, p.Host, p.Port)
	default:

	}
	return p.DSN
}
