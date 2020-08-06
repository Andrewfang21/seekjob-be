package adzuna

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"seekjob/config"
	"seekjob/models"
	"seekjob/utils"
)

// Handler defines operations of scraper handler for Adzuna API
type Handler interface {
	ScrapeJobs()
	getJobs(category, country string, currentPage int) ([]models.Job, error)
	adzunaJobAdapter(rawJob adzunaResult, country string) models.Job
}

type handler struct {
	jobOrmer         models.JobOrmer
	adzunaScraperCfg config.AdzunaScraperCfg
}

// NewAdzunaScraperHandler returns an instance of adzuna scraper handler
func NewAdzunaScraperHandler(
	jobOrmer models.JobOrmer,
	adzunaScraperCfg config.AdzunaScraperCfg) Handler {
	return &handler{
		jobOrmer:         jobOrmer,
		adzunaScraperCfg: adzunaScraperCfg,
	}
}

func (h *handler) ScrapeJobs() {
	// TODO: Use go routine
	for country := range utils.ADZUNA_COUNTRIES_CODE_MAP {
		for category := range utils.ADZUNA_CATEGORIES {
			// Scrape at most 100 pages
			for page := 1; page < 100; page++ {
				jobs, err := h.getJobs(category, country, page)
				if err != nil {
					log.Println(err)
					continue
				}

				if len(jobs) == 0 {
					break
				}

				for _, job := range jobs {
					err := h.jobOrmer.Upsert(job)
					if err != nil {
						log.Printf("[ERROR] Error upsert job in Adzuna %+v: %s", job, err)
						continue
					}
					fmt.Printf("%+v\n", job)
				}
			}
		}
	}
}

func (h *handler) getJobs(category, country string, currentPage int) ([]models.Job, error) {
	r := newAdzunaRequest(h.adzunaScraperCfg, currentPage, country, category)
	endpoints := r.constructEndpoints()
	log.Println(endpoints)

	resp, err := http.Get(endpoints)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error scraping Adzuna with category=%s country=%s page=%d: %s",
			category, country, currentPage, err,
		)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error reading response body: %s", err)
	}

	res := &adzunaResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error unmarshal body: %s", err)
	}

	var jobs []models.Job
	if len(res.Results) == 0 {
		return jobs, nil
	}

	for _, result := range res.Results {
		job := h.adzunaJobAdapter(result, country)
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (h *handler) adzunaJobAdapter(rawJob adzunaResult, location string) models.Job {
	return models.Job{
		ID:          rawJob.ID,
		URL:         rawJob.URL,
		Title:       rawJob.Title,
		Company:     rawJob.Company.Name,
		Description: rawJob.Descripton,
		Category:    rawJob.Category.Label,
		Country:     utils.COUNTRIES_STRING_MAP[utils.ADZUNA_COUNTRIES_CODE_MAP[location]],
		Type:        rawJob.Type,
		PostedAt:    rawJob.PostedAt.Unix(),
		Source:      "ADZ",
	}
}
