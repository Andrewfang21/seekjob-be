package responses

import "seekjob/models"

// JobStatistics model
type JobStatistics struct {
	Results []*JobStatistic `json:"results"`
}

// JobStatistic model
type JobStatistic struct {
	Source     string            `json:"source"`
	Categories []*models.JobInfo `json:"categories"`
	Countries  []*models.JobInfo `json:"countries"`
}

func NewJobStatistics(statistics []*JobStatistic) *JobStatistics {
	return &JobStatistics{Results: statistics}
}
