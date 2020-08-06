package adzuna

import (
	"seekjob/config"
	"seekjob/utils"
	"strconv"
)

const API_BASE_URL = "https://api.adzuna.com"
const API_VERSION = "v1"
const RESULTS_PER_PAGE = "50"

type adzunaRequestable interface {
	constructEndpoints() string
}

type adzunaRequest struct {
	applicationID  string
	applicationKey string
	currentPage    int
	country        string
	category       string
}

func newAdzunaRequest(
	cfg config.AdzunaScraperCfg,
	currentPage int,
	country, category string) adzunaRequestable {
	return &adzunaRequest{
		applicationID:  cfg.ApplicationID,
		applicationKey: cfg.ApplicationKey,
		currentPage:    currentPage,
		country:        country,
		category:       category,
	}
}

/*
	Params: @required app_id
			@required app_key
			@required page
			@optional results_per_age
			@optional category

	By default, Adzuna API returns at most 50 results per page
*/
func (a *adzunaRequest) constructEndpoints() string {
	pageInString := strconv.Itoa(a.currentPage)
	endpoint :=
		utils.ConstructAPIPath(API_BASE_URL, API_VERSION, "api", "jobs", a.country, "search", pageInString) +
			"?" +
			utils.ConstructAPIQuery("app_id", a.applicationID, "app_key", a.applicationKey, "category", a.category, "results_per_page", RESULTS_PER_PAGE)

	return endpoint
}
