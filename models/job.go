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
		SELECT * FROM jobs
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
	queryString := `
		SELECT * FROM jobs
		WHERE
			description ILIKE $1 AND
			category ILIKE $2 AND
			country ILIKE $3 AND
			source ILIKE $4
		OFFSET $5 ROWS
		LIMIT $6
	`
	queryResult, err := j.ormer.Query(queryString,
		query,
		category,
		country,
		source,
		offset,
		limit,
	)
	if err != nil {
		return nil, err
	}
	if queryResult == nil {
		return nil, nil
	}
	defer queryResult.Close()

	var jobs []*Job
	for queryResult.Next() {
		var job Job
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
		jobs = append(jobs, &job)
	}
	return jobs, err
}

func (j *jobOrmer) Upsert(job Job) error {
	queryString := `
		INSERT INTO jobs(id, url, title, company, description, category, country, type, time, source)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT(id)
		DO UPDATE
			SET url=excluded.url,
				title=excluded.title,
				company=excluded.company,
				description=excluded.description,
				category=excluded.category,
				country=excluded.country,
				type=excluded.type,
				time=excluded.time,
				source=excluded.source
	`
	_, err := j.ormer.Exec(queryString,
		job.ID,
		job.URL,
		job.Title,
		job.Company,
		job.Description,
		job.Category,
		job.Country,
		job.Type,
		job.PostedAt,
		job.Source,
	)
	return err
}
