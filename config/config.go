package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var c *Config

func Get() *Config {
	return c
}

type Config struct {
	Env            string `mapstructure:"ENV"`
	ServiceName    string `mapstructure:"SERVICE_NAME"`
	AppURL         string `mapstructure:"APP_URL"`
	Port           string `mapstructure:"PORT"`
	LogLevel       string `mapstructure:"LOG_LEVEL"`
	LogRequestBody bool   `mapstructure:"LOG_REQUEST_BODY"`

	// KAFKA
	KafkaBrokers string `mapstructure:"KAFKA_BROKERS"`
}

func SetConfig(path string) {
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("config not found: %s", err.Error()))
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("SetConfig | could not parse config: %v", err)
	}

	viper.WatchConfig()
}
