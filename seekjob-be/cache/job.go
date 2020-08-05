package cache

import (
	"seekjob/models"
	"seekjob/redis"
)

type JobHandler interface {
	GetSearchedJobs(query string) ([]*models.Job, error)
	SetSearchedJobs(query string, jobs []*models.Job) error
}

type jobHandler struct {
	redisHandler redis.Handler
}

func NewJobCacheHandler(redisHandler redis.Handler) JobHandler {
	return &jobHandler{redisHandler: redisHandler}
}

func (j *jobHandler) GetSearchedJobs(query string) ([]*models.Job, error) {
	return nil, nil
}

func (j *jobHandler) SetSearchedJobs(query string, job []*models.Job) error {
	return nil
}

func redisKey(id, location, category string) string {
	return "jobs:query:" + id + "location:" + location + "category:" + category
}
