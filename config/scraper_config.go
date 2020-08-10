package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var ScraperCfg scraperConfig

type scraperConfig struct {
	Adzuna     AdzunaScraperCfg     `mapstructure:"adzuna"`
	GithubJobs GithubJobsScraperCfg `mapstructure:"github"`
	Jooble     JoobleScraperCfg     `mapstructure:"jooble"`
	Reed       ReedScraperCfg       `mapstructure:"reed"`
	Remotive   RemotiveScraperCfg   `mapstructure:"remotive"`
	TheMuse    TheMuseScraperCfg    `mapstructure:"themuse"`
}

type AdzunaScraperCfg struct {
	ApplicationID  string `mapstructure:"app_id"`
	ApplicationKey string `mapstructure:"app_key"`
}

type GithubJobsScraperCfg struct{}

type JoobleScraperCfg struct {
	ApiKey string `mapstructure:"api_key"`
}

type ReedScraperCfg struct {
	ApiKey string `mapstructure:"api_key"`
}

type RemotiveScraperCfg struct{}

type TheMuseScraperCfg struct {
	ApiKey string `mapstructure:"api_key"`
}

func init() {
	configName := "scraper_config"
	viper.SetConfigName(configName)
	// viper.AddConfigPath(getAppBasePath())
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[ERROR] Fatal error scraper config file: %s", err)
		return
	}

	if err := viper.Unmarshal(&ScraperCfg); err != nil {
		log.Fatalf("[ERROR] Fatal error unmarshal scraper config: %s", err)
		return
	}

	fmt.Println("HEllo here")
}
