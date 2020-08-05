package config

import (
	"log"

	"github.com/spf13/viper"
)

var ScraperCfg scraperConfig

type scraperConfig struct {
	Adzuna     AdzunaScraperCfg     `mapstructure:"adzuna"`
	GithubJobs GithubJobsScraperCfg `mapstructure:"github"`
	Remotive   RemotiveScraperCfg   `mapstructure:"remotive"`
}

type AdzunaScraperCfg struct {
	ApplicationID  string `mapstructure:"app_id"`
	ApplicationKey string `mapstructure:"app_key"`
}

type GithubJobsScraperCfg struct{}

type RemotiveScraperCfg struct{}

func init() {
	viper.SetConfigFile("scraper_config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[ERROR] Fatal error scraper config file: %s", err)
		return
	}

	if err := viper.Unmarshal(&ScraperCfg); err != nil {
		log.Fatalf("[ERROR] Fatal error unmarshal scraper config: %s", err)
		return
	}
}
