package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config is the project's config
var Config config
var environment string = "dev"

type config struct {
	CorsCfg     corsCfg     `mapstructure:"cors"`
	PostgresCfg postgresCfg `mapstructure:"postgres"`
	RedisCfg    redisCfg    `mapstructure:"redis"`
}

type corsCfg struct {
	AllowedMethods []string `mapstructure:"methods"`
	AllowedOrigins []string `mapstructure:"origins"`
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

	configName := fmt.Sprintf("app.%s", environment)
	viper.SetConfigName(configName)
	viper.AddConfigPath(getAppBasePath())

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[ERROR] Fatal error config file: %s", err)
		return
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("[ERROR] Fatal error unmarshal config: %s", err)
		return
	}
}

func getAppBasePath() string {
	basePath, _ := filepath.Abs(".")
	for filepath.Base(basePath) != "seekjob" {
		basePath = filepath.Dir(basePath)
	}
	return basePath
}
