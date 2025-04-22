package config

import (
	"os"

	"github.com/ividernvi/algohub/internal/pkg/options"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Options *options.Options
}

var (
	defaultConfig          *Config
	ALGOHUB_MYSQL_HOSTNAME = os.Getenv("ALGOHUB_MYSQL_HOSTNAME")
	ALGOHUB_MYSQL_PORT     = os.Getenv("ALGOHUB_MYSQL_PORT")
	ALGOHUB_MYSQL_USERNAME = os.Getenv("ALGOHUB_MYSQL_USERNAME")
	ALGOHUB_MYSQL_PASSWORD = os.Getenv("ALGOHUB_MYSQL_PASSWORD")
	ALGOHUB_MYSQL_DATABASE = os.Getenv("ALGOHUB_MYSQL_DATABASE")
)

var (
	ALGOHUB_MINIO_ENDPOINT          = os.Getenv("ALGOHUB_MINIO_ENDPOINT")
	ALGOHUB_MINIO_ACCESS_KEY_ID     = os.Getenv("ALGOHUB_MINIO_ACCESS_KEY_ID")
	ALGOHUB_MINIO_SECRET_ACCESS_KEY = os.Getenv("ALGOHUB_MINIO_SECRET_ACCESS_KEY")
	ALGOHUB_MINIO_USE_SSL           = os.Getenv("ALGOHUB_MINIO_USE_SSL")
	ALGOHUB_MINIO_BUCKET_NAME       = os.Getenv("ALGOHUB_MINIO_BUCKET_NAME")
	ALGOHUB_JUDGE_RPC_ENDPOINT      = os.Getenv("ALGOHUB_JUDGE_RPC_ENDPOINT")
)

func newConfig() *Config {
	return &Config{
		Options: options.NewOptions(),
	}
}

func DefaultConfig() *Config {
	if defaultConfig == nil {
		getDefault()
	}
	return defaultConfig
}

func getDefault() {
	checkDefaultEnv()
	defaultConfig = newConfig()
	defaultConfig.Options.MySQLOpts.HostName = ALGOHUB_MYSQL_HOSTNAME
	defaultConfig.Options.MinioOpts.Endpoint = ALGOHUB_MINIO_ENDPOINT
}

func checkDefaultEnv() {
	if ALGOHUB_MYSQL_HOSTNAME == "" {
		logrus.Printf("database hostname is not set, please set ALGOHUB_MYSQL_HOSTNAME")
		os.Exit(1)
	}

	if ALGOHUB_MINIO_ENDPOINT == "" {
		logrus.Printf("minio endpoint is not set, please set ALGOHUB_MINIO_ENDPOINT")
		os.Exit(1)
	}

	if ALGOHUB_JUDGE_RPC_ENDPOINT == "" {
		logrus.Printf("judge rpc endpoint is not set, please set ALGOHUB_JUDGE_RPC_ENDPOINT")
		os.Exit(1)
	}
}
