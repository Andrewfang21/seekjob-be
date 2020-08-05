package github

import "seekjob/models"

type Handler interface {
	ScrapeJobs()
	getJobsByCountry(country string, currentPage, perPage int) ([]*models.Job, error)
	githubJobAdapter() *models.Job
}

type handler struct {
	jobOrmer models.JobOrmer
}

func NewGithubJobsScraperHandler(jobOrmer models.JobOrmer) Handler {
	return &handler{jobOrmer: jobOrmer}
}

func (h *handler) ScrapeJobs() {

}

func (h *handler) getJobsByCountry(country string, currentPage, perPage int) ([]*models.Job, error) {
	return nil, nil
}

func (h *handler) githubJobAdapter() *models.Job {
	return nil
}
