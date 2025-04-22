package options

type Options struct {
	MySQLOpts            *MySQLOptions         `mapstructure:"mysql_opts"`
	InsecureServeOptions *InsecureServeOptions `mapstructure:"insecure_serve_options"`
	SecureServeOptions   *SecureServeOptions   `mapstructure:"secure_serve_options"`
	RPCServeOptions      *RPCServeOptions      `mapstructure:"rpc_serve_options"`
	MinioOpts            *MinioOptions         `mapstructure:"minio_opts"`
}

func NewOptions() *Options {
	return &Options{
		MySQLOpts:            NewMySQLOptions(),
		InsecureServeOptions: NewInsecureServeOptions(),
		SecureServeOptions:   NewSecureServeOptions(),
		RPCServeOptions:      NewRPCServeOptions(),
		MinioOpts:            NewMinioOptions(),
	}
}
