package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (cfg *Config) LoadConfigFromEnv() (*Config, error) {
	viper.SetConfigName("algohub")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	logrus.Info("Using config file: ", viper.AllSettings())

	if err := viper.Unmarshal(cfg.Options); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}

	fmt.Printf("cfg: %#+v", cfg.Options)

	return &Config{cfg.Options}, nil
}
