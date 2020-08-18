package controllers

import (
	"net/http"
	"reflect"
	"seekjob/cache"
	"seekjob/controllers/requests"
	"seekjob/controllers/responses"
	"seekjob/models"
	"time"

	"github.com/mitchellh/mapstructure"
)

// JobController defines the operations of jobController
type JobController interface {
	GetJob(id string) (*responses.JobResponse, int, *responses.ErrorResponse)
	GetJobs(params requests.JobRequest) (*responses.JobsResponse, int, *responses.ErrorResponse)
	GetJobsStatistics() (*responses.JobStatistics, int, *responses.ErrorResponse)
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
	Headers: -
	Route: /api/jobs/id/:id
	Params: -
	Returns job with specified ID
*/
func (j *jobController) GetJob(id string) (*responses.JobResponse, int, *responses.ErrorResponse) {
	cache, _ := j.jobCacheHandler.GetJob(id)
	if cache != nil {
		return responses.NewJobResponse(cache), http.StatusOK, nil
	}

	job, err := j.jobOrmer.Get(id)
	if err != nil {
		return nil, http.StatusInternalServerError, responses.NewErrorResponse(err.Error())
	}
	if job == nil {
		return nil, http.StatusNoContent, responses.NewErrorResponse("Job with the corresponding ID is not found")
	}

	expiryDuration := time.Duration(time.Hour * 2)
	j.jobCacheHandler.SetJob(job, expiryDuration)

	return responses.NewJobResponse(job), 200, nil
}

/*
	Method: [GET]
	Headers: -
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

	offset := (params.PageNo - 1) * params.PerPage
	jobs, err := j.jobOrmer.GetAll(
		params.Query,
		params.Category,
		params.Country,
		params.Source,
		offset,
		params.PerPage,
	)
	if err != nil {
		return nil, http.StatusInternalServerError, responses.NewErrorResponse(err.Error())
	}
	if jobs == nil {
		return nil, http.StatusNoContent, responses.NewErrorResponse("Jobs not found")
	}

	return responses.NewJobsResponse(jobs), http.StatusOK, nil
}

/*
	Method: [GET]
	Headers: -
	Route: /api/jobs/stats
	Params: -
	Returns list of jobs grouped by categories and the job counts
*/
func (j *jobController) GetJobsStatistics() (*responses.JobStatistics, int, *responses.ErrorResponse) {
	sources, err := j.jobOrmer.GetSources()
	if err != nil {
		return nil, http.StatusInternalServerError, responses.NewErrorResponse(err.Error())
	}

	var statistics []*responses.JobStatistic
	for _, source := range sources {
		var categories, countries []*models.JobInfo
		expiryDuration := time.Duration(time.Hour * 2)

		categories, _ = j.jobCacheHandler.GetCategories(source)
		if categories == nil {
			categories, err = j.jobOrmer.GetCategories(source)
			if err != nil {
				return nil, http.StatusInternalServerError, responses.NewErrorResponse(err.Error())
			}

			j.jobCacheHandler.SetCategories(source, categories, expiryDuration)
		}

		countries, _ = j.jobCacheHandler.GetCountries(source)
		if countries == nil {
			countries, err = j.jobOrmer.GetCountries(source)
			if err != nil {
				return nil, http.StatusInternalServerError, responses.NewErrorResponse(err.Error())
			}

			j.jobCacheHandler.SetCountries(source, countries, expiryDuration)
		}

		statistic := responses.NewJobStatistic(source, categories, countries)
		statistics = append(statistics, statistic)
	}

	return responses.NewJobStatistics(statistics), http.StatusOK, nil
}

func assignDefaultParams(params requests.JobRequest) requests.JobRequest {
	queries := make(map[string]interface{})
	mapstructure.Decode(params, &queries)

	for k, v := range queries {
		if reflect.TypeOf(v).Kind() == reflect.String {
			if v == "" || k == "Query" {
				queries[k] = "%" + v.(string) + "%"
			}
		}
	}

	mapstructure.Decode(queries, &params)
	return params
}
