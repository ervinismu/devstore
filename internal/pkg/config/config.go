package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBConnection string `mapstructure:"DB_CONNECTION"`
	ServerPort   string `mapstructure:"SERVER_PORT"`
	LogLevel     string `mapstructure:"LOG_LEVEL"`
}

func LoadConfig(fileConfigPath string) (Config, error) {
	var config Config

	viper.AddConfigPath(fileConfigPath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Error(fmt.Errorf("error LoadConfig : %w", err))
	}
	return config, nil
}
