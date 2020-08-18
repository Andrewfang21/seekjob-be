package reed

import (
	"encoding/json"
	"fmt"
	"log"
	"seekjob/config"
	"seekjob/models"
	"seekjob/utils"
	"strconv"
	"time"
)

const RESULTS_PER_PAGE int = 100

// Handler defines operations of scraper handler for Reed API
type Handler interface {
	ScrapeJobs()
	getJobs(country, category string, offset int) ([]models.Job, error)
	reedJobAdapter(rawJob reedResult, category, country string) models.Job
}

type handler struct {
	jobOrmer       models.JobOrmer
	reedScraperCfg config.ReedScraperCfg
}

// NewReedScraperHandler returns an instance of reed scraper handler
func NewReedScraperHandler(
	jobOrmer models.JobOrmer,
	reedScraperCfg config.ReedScraperCfg) Handler {
	return &handler{
		jobOrmer:       jobOrmer,
		reedScraperCfg: reedScraperCfg,
	}
}

func (h *handler) ScrapeJobs() {
	// TODO: Use go routine
	countries, err := utils.GetCountries("REED")
	if err != nil {
		log.Println(err)
	}
	categories, err := utils.GetCategories("REED")
	if err != nil {
		log.Println(err)
	}

	for _, country := range countries {
		for _, category := range categories {
			// Scrape at most 100 pages
			for page := 1; page < 100; page++ {
				offset := (page - 1) * RESULTS_PER_PAGE
				jobs, err := h.getJobs(country, category, offset)
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
						log.Printf("[ERROR] Error upsert job in Reed %+v: %s", job, err)
						continue
					}
				}
			}
		}
	}

}

func (h *handler) getJobs(country, category string, offset int) ([]models.Job, error) {
	r := newReedRequest(h.reedScraperCfg.APIKey, country, category, offset)
	body, err := r.callEndpoint("GET")
	if err != nil {
		return nil, err
	}

	resp := &reedResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error unmarshal Reed body: %s", err)
	}

	var jobs []models.Job
	if len(resp.Results) == 0 {
		return jobs, nil
	}

	for _, result := range resp.Results {
		job := h.reedJobAdapter(result, category, country)
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (h *handler) reedJobAdapter(rawJob reedResult, category, country string) models.Job {
	idInString := strconv.Itoa(rawJob.ID)
	timestamp, _ := time.Parse("02/01/2006", rawJob.PostedAt)

	return models.Job{
		ID:          idInString,
		URL:         rawJob.URL,
		Title:       rawJob.Title,
		Company:     rawJob.Company,
		Description: rawJob.Description,
		Category:    category,
		Country:     country,
		PostedAt:    timestamp.Unix(),
		Type:        "",
		Source:      "REE",
	}
}
