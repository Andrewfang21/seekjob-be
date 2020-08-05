package adzuna

import (
	"fmt"
	"seekjob/config"
)

const API_BASE_URL string = "https://api.adzuna.com"
const API_VERSION string = "v1"

type adzunaRequestable interface {
	constructEndpoints() string
}

type adzunaRequest struct {
	applicationID  string
	applicationKey string
	currentPage    int
	perPage        int
	country        string
}

func NewAdzunaRequest(
	cfg config.AdzunaScraperCfg,
	currentPage, perPage int,
	country string) adzunaRequestable {
	return &adzunaRequest{
		applicationID:  cfg.ApplicationID,
		applicationKey: cfg.ApplicationKey,
		currentPage:    currentPage,
		perPage:        perPage,
		country:        country,
	}
}

func (a *adzunaRequest) constructEndpoints() string {
	apiEndpoint := fmt.Sprintf("%s/%s/api/jobs/%s/search/%d?app_id=%s&app_key=%s&results_per_page=%d",
		API_BASE_URL, API_VERSION,
		a.country,
		a.currentPage,
		a.applicationID, a.applicationKey, a.perPage,
	)
	return apiEndpoint
}
