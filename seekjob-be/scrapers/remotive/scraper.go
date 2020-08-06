package remotive

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"seekjob/models"
	"seekjob/utils"
	"strconv"
	"time"
)

// Handler defines operations of scraper handler for Remotive API
type Handler interface {
	ScrapeJobs()
	getJobsByCategory(category string) ([]models.Job, error)
	remotiveJobAdapter(rawJob remotiveJob) models.Job
}

type handler struct {
	jobOrmer models.JobOrmer
}

// NewRemotiveScraperHandler returns an instance of remotive scraper handler
func NewRemotiveScraperHandler(jobOrmer models.JobOrmer) Handler {
	return &handler{jobOrmer: jobOrmer}
}

func (h *handler) ScrapeJobs() {
	// TODO: Use go routine
	// By default, Remotive does not support pagination
	for category := range utils.REMOTIVE_JOBS_CATEGORIES {
		jobs, err := h.getJobsByCategory(category)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(jobs) == 0 {
			continue
		}

		for _, job := range jobs {
			err := h.jobOrmer.Upsert(job)
			if err != nil {
				log.Printf("[ERROR] Error upsert job in Remotive %+v: %s", job, err)
				continue
			}
			fmt.Printf("%+v\n", job.PostedAt)
		}
	}
}

func (h *handler) getJobsByCategory(category string) ([]models.Job, error) {
	r := newRemotiveRequest(category)
	endpoints := r.constructEndpoints()
	log.Println(endpoints)

	resp, err := http.Get(endpoints)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error scraping Remotive with category=%s: %s",
			category, err,
		)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error reading response body: %s", err)
	}

	var results remotiveResponse
	err = json.Unmarshal(body, &results)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error unmarshal body: %s", err)
	}

	var jobs []models.Job
	if len(results.Jobs) == 0 {
		return jobs, nil
	}

	for _, result := range results.Jobs {
		job := h.remotiveJobAdapter(result)
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (h *handler) remotiveJobAdapter(rawJob remotiveJob) models.Job {
	idInString := strconv.Itoa(rawJob.ID)
	timestamp, _ := time.Parse(time.RFC3339, rawJob.PostedAt+"Z")
	return models.Job{
		ID:          idInString,
		URL:         rawJob.URL,
		Title:       rawJob.Title,
		Company:     rawJob.Company,
		Description: rawJob.Description,
		Category:    rawJob.Category,
		Country:     "REM",
		PostedAt:    timestamp.Unix(),
		Type:        rawJob.Type,
		Source:      "REM",
	}
}
