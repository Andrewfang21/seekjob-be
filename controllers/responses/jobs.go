package responses

import "seekjob/models"

// JobsResponse model
type JobsResponse struct {
	Count int          `json:"job_count"`
	Jobs  []models.Job `json:"jobs"`
}

// NewJobsResponse returns the json response for jobs
func NewJobsResponse(jobs []models.Job) *JobsResponse {
	return &JobsResponse{
		Count: len(jobs),
		Jobs:  jobs,
	}
}
