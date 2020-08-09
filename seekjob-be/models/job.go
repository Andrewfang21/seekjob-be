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
	Get(id int) (*Job, error)
	GetAll(query string, offset, limit int) ([]*Job, error)
	GetByCategory(query, category string, offset, limit int) ([]*Job, error)
	GetByLocation(query, location string, offset, limit int) ([]*Job, error)
	Upsert(job Job) error
}

type jobOrmer struct {
	ormer *sql.DB
}

// NewJobOrmer returns an instance of jobOrmer
func NewJobOrmer(ormer *sql.DB) JobOrmer {
	return &jobOrmer{ormer: ormer}
}

func (j *jobOrmer) Get(id int) (*Job, error) {
	return nil, nil
}

func (j *jobOrmer) GetAll(query string, offset, limit int) ([]*Job, error) {
	return nil, nil
}

func (j *jobOrmer) GetByCategory(query, category string, offset, limit int) ([]*Job, error) {
	return nil, nil
}

func (j *jobOrmer) GetByLocation(query, location string, offset, limit int) ([]*Job, error) {
	return nil, nil
}

func (j *jobOrmer) Upsert(job Job) error {
	return nil
}
