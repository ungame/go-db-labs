package db

import (
	"flag"
	"fmt"
)

const (
	DefaultMysqlHost   = "localhost"
	DefaultMysqlPort   = 3306
	DefaultMysqlUser   = "root"
	DefaultMysqlPass   = "root"
	DefaultMysqlDbName = "lab"
)

type MysqlConfig interface {
	Dsn() string
}

type config struct {
	mysqlHost   string
	mysqlPort   int
	mysqlUser   string
	mysqlPass   string
	mysqlDbName string
}

func NewConfigFromFlags(setter *flag.FlagSet) MysqlConfig {
	var cfg config
	setter.StringVar(&cfg.mysqlHost, "mysql_host", DefaultMysqlHost, "set mysql host")
	setter.IntVar(&cfg.mysqlPort, "mysql_port", DefaultMysqlPort, "set mysql port")
	setter.StringVar(&cfg.mysqlUser, "mysql_user", DefaultMysqlUser, "set mysql username")
	setter.StringVar(&cfg.mysqlPass, "mysql_pass", DefaultMysqlPass, "set mysql password")
	setter.StringVar(&cfg.mysqlDbName, "mysql_dbname", DefaultMysqlDbName, "set mysql dbname")
	return &cfg
}

func NewDefaultConfig() MysqlConfig {
	return &config{
		mysqlHost:   DefaultMysqlHost,
		mysqlPort:   DefaultMysqlPort,
		mysqlUser:   DefaultMysqlUser,
		mysqlPass:   DefaultMysqlPass,
		mysqlDbName: DefaultMysqlDbName,
	}
}

func (c *config) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.mysqlUser, c.mysqlPass, c.mysqlHost, c.mysqlPort, c.mysqlDbName)
}
