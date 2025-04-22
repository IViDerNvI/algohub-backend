package options

import (
	"fmt"
)

type RPCServeOptions struct {
	BindAddr string `json:"address" mapstructure:"bind-addr"`
	BindPort int    `json:"port" mapstructure:"bind-port"`
}

func NewRPCServeOptions() *RPCServeOptions {
	return &RPCServeOptions{
		BindAddr: "0.0.0.0",
		BindPort: 50051,
	}
}

func (o *RPCServeOptions) Address() string {
	return o.BindAddr
}

func (o *RPCServeOptions) Port() string {
	return fmt.Sprintf(":%d", o.BindPort)
}
