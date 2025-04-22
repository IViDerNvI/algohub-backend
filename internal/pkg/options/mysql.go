package options

import (
	"strconv"
	"time"
)

type MySQLOptions struct {
	HostName              string        `json:"host-name" mapstructure:"host-name"`
	BindPort              int           `json:"bind-port" mapstructure:"bind-port"`
	User                  string        `json:"username" mapstructure:"username"`
	Passwd                string        `json:"assword" mapstructure:"password"`
	DatabaseNo            string        `json:"database" mapstructure:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections" mapstructure:"max-idle-connections"`
	MaxOpenConnections    int           `json:"max-open-connections" mapstructure:"max-open-connections"`
	MaxConnectionLifetime time.Duration `json:"max-connection-lifetime" mapstructure:"max-connection-lifetime"`
}

func (o *MySQLOptions) Host() string {
	return o.HostName
}

func (o *MySQLOptions) Port() string {
	return strconv.Itoa(o.BindPort)
}

func (o *MySQLOptions) Username() string {
	return o.User
}

func (o *MySQLOptions) Password() string {
	return o.Passwd
}

func (o *MySQLOptions) Database() string {
	return o.DatabaseNo
}

func (o *MySQLOptions) MaxIdleConns() int {
	return o.MaxIdleConnections
}

func (o *MySQLOptions) MaxOpenConns() int {
	return o.MaxOpenConnections
}

func (o *MySQLOptions) MaxLifetime() time.Duration {
	return o.MaxConnectionLifetime
}

func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		HostName:              "127.0.0.1",
		BindPort:              3306,
		User:                  "root",
		Passwd:                "root",
		DatabaseNo:            "algohub",
		MaxIdleConnections:    10,
		MaxOpenConnections:    100,
		MaxConnectionLifetime: 10 * time.Minute,
	}
}
