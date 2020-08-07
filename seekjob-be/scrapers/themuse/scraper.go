package themuse

import (
	"encoding/json"
	"fmt"
	"log"
	"seekjob/config"
	"seekjob/models"
	"seekjob/utils"
	"strconv"
)

// Handler defines operations of scraper handler for TheMuse API
type Handler interface {
	ScrapeJobs()
	getJobs(category string, currentPage int) ([]models.Job, error)
	theMuseAdapter(rawJob theMuseResult, category string) models.Job
}

type handler struct {
	jobOrmer          models.JobOrmer
	theMuseScraperCfg config.TheMuseScraperCfg
}

// NewTheMuseScraperHandler returns an instance of themuse scraper handler
func NewTheMuseScraperHandler(
	jobOrmer models.JobOrmer,
	theMuseScraperCfg config.TheMuseScraperCfg) Handler {
	return &handler{
		jobOrmer:          jobOrmer,
		theMuseScraperCfg: theMuseScraperCfg,
	}
}

func (h *handler) ScrapeJobs() {
	for _, category := range utils.THE_MUSE_JOBS_CATEGORIES {
		// Scrape at most 500 pages
		for page := 1; page < 2; page++ {
			jobs, err := h.getJobs(category, page)
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
					log.Printf("[ERROR] Error upsert job in TheMuse %+v: %s", job, err)
					continue
				}
			}
		}
	}
}

func (h *handler) getJobs(category string, currentPage int) ([]models.Job, error) {
	r := newTheMuseRequest(h.theMuseScraperCfg.ApiKey, category, currentPage)
	body, err := r.callEndpoint("GET")
	if err != nil {
		return nil, err
	}

	resp := &theMuseResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error unmarshal body: %s", err)
	}

	var jobs []models.Job
	if len(resp.Results) == 0 {
		return jobs, nil
	}

	for _, result := range resp.Results {
		job := h.theMuseAdapter(result, category)
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (h *handler) theMuseAdapter(rawJob theMuseResult, category string) models.Job {
	idInString := strconv.Itoa(rawJob.ID)

	return models.Job{
		ID:          idInString,
		URL:         rawJob.URL.LandingPage,
		Title:       rawJob.Title,
		Company:     rawJob.Company.Name,
		Description: rawJob.Description,
		Category:    category,
		Country:     "", // TODO: TheMuse jobs' location is hard to map
		PostedAt:    rawJob.PostedAt.Unix(),
		Type:        "",
		Source:      "TMS",
	}
}
