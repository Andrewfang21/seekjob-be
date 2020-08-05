package github

import "fmt"

const API_BASE_URL string = "https://jobs.github.com/positions.json"

type githubJobsRequestable interface {
	constructEndpoints() string
}

type githubJobsRequest struct {
	currentPage int
	country     string
}

func NewGithubJobsRequest(currentPage int, country string) githubJobsRequestable {
	return &githubJobsRequest{
		currentPage: currentPage,
		country:     country,
	}
}

func (g *githubJobsRequest) constructEndpoints() string {
	apiEndpoint := fmt.Sprintf("%s?location=%s&page=%d",
		API_BASE_URL,
		g.country, g.currentPage,
	)
	return apiEndpoint
}
