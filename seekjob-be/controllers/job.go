package controllers

import (
	"fmt"
	"net/http"
	"seekjob/cache"
	"seekjob/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// JobController defines the operations of jobController
type JobController interface {
	GetJob(*gin.Context)
	GetJobs(*gin.Context)
	GetJobsByCategory(*gin.Context)
	GetJobsByLocation(*gin.Context)
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
func (j *jobController) GetJob(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id should be an integer",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Routed inside GetJob() with id: %d", id),
	})
}

/*
	Method: [GET]
	Route: /api/jobs
	Params: @required page_no `int`
			@required per_page `int`
			q `string`
	Returns list of jobs
*/
func (j *jobController) GetJobs(c *gin.Context) {
	query := c.DefaultQuery("q", "*")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Routed inside GetJobs() with query: %s", query),
	})
}

/*
	Method: [GET]
	Route: /api/jobs/category/:category
	Params: @required page_no `int`
			@required per_page `int`
			q `string`
	Returns list of jobs with specified category
*/
func (j *jobController) GetJobsByCategory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Routed inside GetJobsByCategory()",
	})
}

/*
	Method: [GET]
	Route: /api/jobs/location/:location
	Params: @required page_no `int`
			@required per_page `int`
			q `string`
	Returns list of jobs with specified country
*/
func (j *jobController) GetJobsByLocation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Routed inside GetJobsByLocation()",
	})
}
