package responses

import "seekjob/models"

// JobResponse model
type JobResponse struct {
	Job *models.Job `json:"job"`
}

// NewJobResponse returns the json response for job
func NewJobResponse(job *models.Job) *JobResponse {
	return &JobResponse{Job: job}
}
