package options

import (
	"fmt"
)

type InsecureServeOptions struct {
	BindAddr string `json:"BindAddr" mapstructure:"bind-addr"`
	BindPort int    `json:"BindPort" mapstructure:"bind-port"`
}

func NewInsecureServeOptions() *InsecureServeOptions {
	return &InsecureServeOptions{
		BindAddr: "0.0.0.0",
		BindPort: 8080,
	}
}

func (o *InsecureServeOptions) Addr() string {
	return o.BindAddr
}

func (o *InsecureServeOptions) Port() string {
	return fmt.Sprintf(":%d", o.BindPort)
}
