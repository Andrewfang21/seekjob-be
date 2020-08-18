package cache

import (
	"seekjob/models"
	"seekjob/redis"
	"time"
)

// JobHandler defines the operations of jobHandler
type JobHandler interface {
	GetJob(id string) (*models.Job, error)
	GetCountries(source string) ([]*models.JobInfo, error)
	GetCategories(source string) ([]*models.JobInfo, error)
	SetJob(job *models.Job, expiryDuration time.Duration) error
	SetCountries(source string, countries []*models.JobInfo, expiryDuration time.Duration) error
	SetCategories(source string, categories []*models.JobInfo, expiryDuration time.Duration) error
}

type jobHandler struct {
	redisHandler redis.Handler
}

// NewJobCacheHandler returns job cache handler
func NewJobCacheHandler(redisHandler redis.Handler) JobHandler {
	return &jobHandler{redisHandler: redisHandler}
}

func (j *jobHandler) GetJob(id string) (*models.Job, error) {
	var job models.Job

	ok, err := j.redisHandler.Get(jobsRedisKey(id), &job)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return &job, nil
}

func (j *jobHandler) GetCountries(source string) ([]*models.JobInfo, error) {
	var countries []*models.JobInfo

	ok, err := j.redisHandler.Get(countriesRedisKey(source), &countries)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return countries, nil
}

func (j *jobHandler) GetCategories(source string) ([]*models.JobInfo, error) {
	var categories []*models.JobInfo

	ok, err := j.redisHandler.Get(categoriesRedisKey(source), &categories)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return categories, nil
}

func (j *jobHandler) SetJob(job *models.Job, expiryDuration time.Duration) error {
	return j.redisHandler.SetWithExpiry(jobsRedisKey(job.ID), job, expiryDuration)
}

func (j *jobHandler) SetCountries(source string, countries []*models.JobInfo, expiryDuration time.Duration) error {
	return j.redisHandler.SetWithExpiry(countriesRedisKey(source), countries, expiryDuration)
}

func (j *jobHandler) SetCategories(source string, categories []*models.JobInfo, expiryDuration time.Duration) error {
	return j.redisHandler.SetWithExpiry(categoriesRedisKey(source), categories, expiryDuration)
}

func jobsRedisKey(id string) string {
	return "jobs:id:" + id
}

func countriesRedisKey(source string) string {
	return "countries:source:" + source
}

func categoriesRedisKey(source string) string {
	return "categories:source:" + source
}
