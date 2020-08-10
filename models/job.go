package models

import (
	"database/sql"
)

// Job model
type Job struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Company     string `json:"company"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Country     string `json:"country"`
	Type        string `json:"type"`
	PostedAt    int64  `json:"time"`
	Source      string `json:"source"`
}

// JobOrmer defines the operations of jobOrmer
type JobOrmer interface {
	Get(id string) (*Job, error)
	GetAll(query, category, country, source string, offset, limit int) ([]*Job, error)
	Upsert(job Job) error
}

type jobOrmer struct {
	ormer *sql.DB
}

// NewJobOrmer returns an instance of jobOrmer
func NewJobOrmer(ormer *sql.DB) JobOrmer {
	return &jobOrmer{ormer: ormer}
}

func (j *jobOrmer) Get(id string) (*Job, error) {
	queryString := `
		SELECT *
		FROM jobs
		WHERE
			id = $1
	`
	queryResult, err := j.ormer.Query(queryString, id)

	if err != nil {
		return nil, err
	}
	if queryResult == nil {
		return nil, nil
	}

	defer queryResult.Close()
	var job Job
	isExists := queryResult.Next()
	if !isExists {
		return nil, nil
	}
	err = queryResult.Scan(
		&job.ID,
		&job.URL,
		&job.Title,
		&job.Company,
		&job.Description,
		&job.Category,
		&job.Country,
		&job.Type,
		&job.PostedAt,
		&job.Source,
	)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (j *jobOrmer) GetAll(query, category, country, source string, offset, limit int) ([]*Job, error) {
	return nil, nil
}

func (j *jobOrmer) Upsert(job Job) error {
	return nil
}
