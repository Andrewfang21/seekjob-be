package github

import (
	"encoding/json"
	"fmt"
	"log"
	"seekjob/models"
	"seekjob/utils"
	"time"
)

// Handler defines operations of scraper handler for GithubJobs API
type Handler interface {
	ScrapeJobs()
	getJobsByCountry(country string, currentPage int) ([]models.Job, error)
	githubJobAdapter(rawJob githubJobsResponse, country string) models.Job
}

type handler struct {
	jobOrmer models.JobOrmer
}

// NewGithubJobsScraperHandler returns an instance of github jobs scraper handler
func NewGithubJobsScraperHandler(jobOrmer models.JobOrmer) Handler {
	return &handler{jobOrmer: jobOrmer}
}

func (h *handler) ScrapeJobs() {
	// TODO: Use go routine
	countries, err := utils.GetCountries("GITHUB")
	if err != nil {
		log.Println(err)
	}
	for _, country := range countries {
		// Scrape at most 100 pages
		for page := 0; page < 100; page++ {
			jobs, err := h.getJobsByCountry(country, page)
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
					log.Printf("[ERROR] Error upsert job in GithubJob %+v: %s", job, err)
					continue
				}
			}
		}
	}
}

func (h *handler) getJobsByCountry(country string, currentPage int) ([]models.Job, error) {
	r := newGithubJobsRequest(currentPage, country)
	body, err := r.callEndpoint("GET")
	if err != nil {
		return nil, err
	}

	var results []githubJobsResponse
	err = json.Unmarshal(body, &results)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error unmarshal body: %s", err)
	}

	var jobs []models.Job
	if len(results) == 0 {
		return jobs, nil
	}

	for _, result := range results {
		job := h.githubJobAdapter(result, country)
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (h *handler) githubJobAdapter(rawJob githubJobsResponse, country string) models.Job {
	timestamp, _ := time.Parse("Mon Jan 02 15:04:05 MST 2006", rawJob.PostedAt)

	return models.Job{
		ID:          rawJob.ID,
		URL:         rawJob.URL,
		Title:       rawJob.Title,
		Company:     rawJob.Company,
		Description: rawJob.Description,
		Category:    "Others",
		Country:     country,
		PostedAt:    timestamp.Unix(),
		Type:        rawJob.Type,
		Source:      "GHJ",
	}
}
