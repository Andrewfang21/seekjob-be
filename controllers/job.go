package controllers

import (
	"reflect"
	"seekjob/cache"
	"seekjob/controllers/requests"
	"seekjob/controllers/responses"
	"seekjob/models"

	"github.com/mitchellh/mapstructure"
)

// JobController defines the operations of jobController
type JobController interface {
	GetJob(id string) (*responses.JobResponse, int, *responses.ErrorResponse)
	GetJobs(params requests.JobRequest) (*responses.JobsResponse, int, *responses.ErrorResponse)
	GetJobsStatistics() (*responses.JobStatistic, int, *responses.ErrorResponse)
}

type jobController struct {
	jobOrmer        models.JobOrmer
	jobCacheHandler cache.JobHandler
}

// NewJobController returns an instance of jobController
func NewJobController(jobOrmer models.JobOrmer, jobCacheHandler cache.JobHandler) JobController {
	return &jobController{
		jobOrmer:        jobOrmer,
		jobCacheHandler: jobCacheHandler,
	}
}

/*
	Method: [GET]
	Route: /api/jobs/id/:id
	Params: -
	Returns job with specified ID
*/
func (j *jobController) GetJob(id string) (*responses.JobResponse, int, *responses.ErrorResponse) {
	return responses.NewJobResponse(models.Job{ID: id}), 200, nil
}

/*
	Method: [GET]
	Route: /api/jobs
	Params: @required page_no `int`
			@required per_page `int`
			q `string`
			category `string`
			country `string`
			source `string`
	Returns list of jobs
*/
func (j *jobController) GetJobs(params requests.JobRequest) (*responses.JobsResponse, int, *responses.ErrorResponse) {
	params = assignDefaultParams(params)
	return nil, 200, nil
}

/*
	Method: [GET]
	Route: /api/jobs/stats
	Params: -
	Returns list of jobs grouped by categories and the job counts
*/
func (j *jobController) GetJobsStatistics() (*responses.JobStatistic, int, *responses.ErrorResponse) {
	return nil, 200, nil
}

func assignDefaultParams(params requests.JobRequest) requests.JobRequest {
	queries := make(map[string]interface{})
	mapstructure.Decode(params, &queries)

	for k, v := range queries {
		if reflect.TypeOf(v).Kind() == reflect.String {
			if v == "" {
				queries[k] = "%"
			}
		}
	}

	mapstructure.Decode(queries, &params)
	return params
}
