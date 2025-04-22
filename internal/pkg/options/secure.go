package options

import (
	"fmt"
)

type SecureServeOptions struct {
	BindAddr    string `json:"bind_addr" mapstructure:"bind-addr"`
	BindPort    int    `json:"bind_port" mapstructure:"bind-port"`
	CertKeyFile CertKey
}

type CertKey struct {
	CertFile string `json:"cert_file" mapstructure:"cert-file"`
	KeyFile  string `json:"key_file" mapstructure:"key-file"`
}

func NewSecureServeOptions() *SecureServeOptions {
	return &SecureServeOptions{
		BindAddr: "0.0.0.0",
		BindPort: 8443,
		CertKeyFile: CertKey{
			CertFile: "./conf/cert/cert.crt",
			KeyFile:  "./conf/cert/key.crt",
		},
	}
}

func (s *SecureServeOptions) Addr() string {
	return s.BindAddr
}

func (s *SecureServeOptions) Port() string {
	return fmt.Sprintf(":%d", s.BindPort)
}

func (s *SecureServeOptions) CertFile() string {
	return s.CertKeyFile.CertFile
}

func (s *SecureServeOptions) KeyFile() string {
	return s.CertKeyFile.KeyFile
}
