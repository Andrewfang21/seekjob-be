package github

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"seekjob/utils"
	"strconv"
)

const API_BASE_URL string = "https://jobs.github.com/positions.json"

type githubJobsRequestable interface {
	callEndpoint(method string) ([]byte, error)
	constructRequestHeaders(req *http.Request)
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
	API Docs: https://jobs.github.com/api
	Params: @optional location

	By default, Github Jobs API returns at most 50 results per page
*/
func (g *githubJobsRequest) callEndpoint(method string) ([]byte, error) {
	pageInString := strconv.Itoa(g.currentPage)
	path := API_BASE_URL
	params := map[string]string{
		"location": g.country,
		"page":     pageInString,
	}

	url := utils.ConstructRequestURL(path, params)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error creating NewRequest: %s", err)
	}
	g.constructRequestHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error scraping GithubJobs with country=%s page=%d: %s",
			g.country, g.currentPage, err,
		)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error reading response body: %s", err)
	}

	return body, nil
}

func (g *githubJobsRequest) constructRequestHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}
