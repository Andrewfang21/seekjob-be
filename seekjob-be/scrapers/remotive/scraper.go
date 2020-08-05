package remotive

import "seekjob/models"

type Handler interface {
	ScrapeJobs()
	getJobsByTag(tag string) ([]*models.Job, error)
	remotiveJobAdapter() *models.Job
}

type handler struct {
	jobOrmer models.JobOrmer
}

func NewRemotiveScraperHandler(jobOrmer models.JobOrmer) Handler {
	return &handler{jobOrmer: jobOrmer}
}

func (h *handler) ScrapeJobs() {

}

func (h *handler) getJobsByTag(tag string) ([]*models.Job, error) {
	return nil, nil
}

func (h *handler) remotiveJobAdapter() *models.Job {
	return nil
}
