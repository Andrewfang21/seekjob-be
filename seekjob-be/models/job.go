package models

import (
	"github.com/jinzhu/gorm"
)

// Job model
type Job struct {
	ID          string `json:"id" gorm:"column:id"`
	URL         string `json:"url" gorm:"column:url"`
	Title       string `json:"title" gorm:"column:title"`
	Company     string `json:"company" gorm:"column:company"`
	Description string `json:"description" gorm:"column:description"`
	Category    string `json:"category" gorm:"column:category"`
	Country     string `json:"country" gorm:"column:country"`
	Type        string `json:"type" gorm:"column:type"`
	PostedAt    int64  `json:"time" gorm:"column:time"`
	Source      string `json:"source" gorm:"column:source"`
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
	ormer *gorm.DB
}

// NewJobOrmer returns an instance of jobOrmer
func NewJobOrmer(ormer *gorm.DB) JobOrmer {
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
