package config

import (
	"log"

	"github.com/spf13/viper"
)

// ScraperCfg exports the scraper config
var ScraperCfg scraperConfig

type scraperConfig struct {
	Adzuna     AdzunaScraperCfg     `mapstructure:"adzuna"`
	GithubJobs GithubJobsScraperCfg `mapstructure:"github"`
	Jooble     JoobleScraperCfg     `mapstructure:"jooble"`
	Reed       ReedScraperCfg       `mapstructure:"reed"`
	Remotive   RemotiveScraperCfg   `mapstructure:"remotive"`
	TheMuse    TheMuseScraperCfg    `mapstructure:"themuse"`
}

// AdzunaScraperCfg is Adzuna API Config
type AdzunaScraperCfg struct {
	ApplicationID  string `mapstructure:"app_id"`
	ApplicationKey string `mapstructure:"app_key"`
}

// GithubJobsScraperCfg is GithubJobs API Config
type GithubJobsScraperCfg struct{}

// JoobleScraperCfg is Jooble API Config
type JoobleScraperCfg struct {
	APIKey string `mapstructure:"api_key"`
}

// ReedScraperCfg is Reed API Config
type ReedScraperCfg struct {
	APIKey string `mapstructure:"api_key"`
}

// RemotiveScraperCfg is Remotive API Config
type RemotiveScraperCfg struct{}

// TheMuseScraperCfg is TheMuse API Config
type TheMuseScraperCfg struct {
	APIKey string `mapstructure:"api_key"`
}

func init() {
	configName := "scraper_config"
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[ERROR] Fatal error scraper config file: %s", err)
		return
	}

	if err := viper.Unmarshal(&ScraperCfg); err != nil {
		log.Fatalf("[ERROR] Fatal error unmarshal scraper config: %s", err)
		return
	}
}
