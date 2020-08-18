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

// NewJobStatistic returns JobStatistic model
func NewJobStatistic(source string, categories, countries []*models.JobInfo) *JobStatistic {
	return &JobStatistic{
		Source:     source,
		Categories: categories,
		Countries:  countries,
	}
}

// NewJobStatistics returns JobStatistics model
func NewJobStatistics(statistics []*JobStatistic) *JobStatistics {
	return &JobStatistics{Results: statistics}
}
