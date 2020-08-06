package github

import (
	"seekjob/utils"
	"strconv"
)

const API_BASE_URL = "https://jobs.github.com/positions.json"

type githubJobsRequestable interface {
	constructEndpoints() string
}

type githubJobsRequest struct {
	currentPage int
	country     string
}

func newGithubJobsRequest(currentPage int, country string) githubJobsRequestable {
	return &githubJobsRequest{
		currentPage: currentPage,
		country:     country,
	}
}

/*
	Params: @optional location
	By default, Github Jobs API returns at most 50 results per page
*/
func (g *githubJobsRequest) constructEndpoints() string {
	pageInString := strconv.Itoa(g.currentPage)
	endpoint := API_BASE_URL + "?" + utils.ConstructAPIQuery("location", g.country, "page", pageInString)
	return endpoint
}
