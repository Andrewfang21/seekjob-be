package adzuna

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"seekjob/config"
	"seekjob/utils"
	"strconv"
)

const API_BASE_URL string = "https://api.adzuna.com"
const API_VERSION string = "v1"
const RESULTS_PER_PAGE string = "50"

type adzunaRequestable interface {
	callEndpoint(method string) ([]byte, error)
	constructRequestHeaders(req *http.Request)
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
	API Docs: https://developer.adzuna.com/activedocs#!/adzuna/categories
	Params: @required app_id
			@required app_key
			@required page
			@optional results_per_age
			@optional category

	By default, Adzuna API returns at most 50 results per page
*/
func (a *adzunaRequest) callEndpoint(method string) ([]byte, error) {
	pageInString := strconv.Itoa(a.currentPage)
	path := utils.ConstructURLPath(
		API_BASE_URL,
		API_VERSION,
		"api", "jobs",
		a.country,
		"search",
		pageInString,
	)
	params := map[string]string{
		"app_id":           a.applicationID,
		"app_key":          a.applicationKey,
		"category":         a.category,
		"results_per_page": RESULTS_PER_PAGE,
	}

	url := utils.ConstructRequestURL(path, params)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error creating NewRequest: %s", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error scraping Adzuna with category=%s country=%s page=%d: %s",
			a.category, a.country, a.currentPage, err,
		)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error reading response body: %s", err)
	}

	return body, nil
}
func (a *adzunaRequest) constructRequestHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}
