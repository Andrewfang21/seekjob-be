package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var Config config
var environment string = "dev"

type config struct {
	PostgresCfg postgresCfg `mapstructure:"postgres"`
	RedisCfg    redisCfg    `mapstructure:"redis"`
}

type postgresCfg struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type redisCfg struct {
	Address              string `mapstructure:"address"`
	Password             string `mapstructure:"password"`
	Database             int    `mapstructure:"database"`
	MaxRetries           int    `mapstructure:"max_retries"`
	DialTimeoutInSeconds int    `mapstructure:"dial_timeout"`
}

func init() {
	env := os.Getenv("APP_ENV")
	if env != "" {
		environment = env
	}

	configFile := fmt.Sprintf("app.%s.yaml", environment)
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[ERROR] Fatal error config file: %s", err)
		return
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("[ERROR] Fatal error unmarshal config: %s", err)
		return
	}
}
