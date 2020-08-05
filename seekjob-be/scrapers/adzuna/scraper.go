package adzuna

import (
	"seekjob/config"
	"seekjob/models"
)

type Handler interface {
	ScrapeJobs()
	getJobsByCountry(country string, currentPage, perPage int) ([]*models.Job, error)
	adzunaJobAdapter() *models.Job
}

type handler struct {
	jobOrmer         models.JobOrmer
	adzunaScraperCfg config.AdzunaScraperCfg
}

func NewAdzunaScraperHandler(
	jobOrmer models.JobOrmer,
	adzunaScraperCfg config.AdzunaScraperCfg) Handler {
	return &handler{
		jobOrmer:         jobOrmer,
		adzunaScraperCfg: adzunaScraperCfg,
	}
}

func (h *handler) ScrapeJobs() {

}

func (h *handler) getJobsByCountry(country string, currentPage, perPage int) ([]*models.Job, error) {
	return nil, nil
}

func (h *handler) adzunaJobAdapter() *models.Job {
	return nil
}
