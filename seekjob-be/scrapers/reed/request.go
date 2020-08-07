package reed

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"seekjob/utils"
	"strconv"
)

const API_BASE_URL string = "https://www.reed.co.uk"
const API_VERSION string = "1.0"
const BASIC_PREFIX string = "Basic "

type reedRequestable interface {
	callEndpoint(method string) ([]byte, error)
	constructRequestHeaders(req *http.Request)
}

type reedRequest struct {
	apiKey   string
	country  string
	category string
	offset   int
}

func newReedRequest(
	apiKey,
	country, category string,
	offset int) reedRequestable {
	return &reedRequest{
		apiKey:   apiKey,
		country:  country,
		category: category,
		offset:   offset,
	}
}

/*
	API Docs: https://www.reed.co.uk/developers/Jobseeker
	Params: @optional locationName
			@optional keywords
			@optional resultsToSkip

	By default, Reed API returns at most 100 results per page
*/
func (r *reedRequest) callEndpoint(method string) ([]byte, error) {
	offset := strconv.Itoa(r.offset)
	path := utils.ConstructUrlPath(
		API_BASE_URL,
		"api",
		API_VERSION,
		"search",
	)
	params := map[string]string{
		"locationName":  r.country,
		"keywords":      r.category,
		"resultsToSkip": offset,
	}

	url := utils.ConstructRequestUrl(path, params)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error creating NewRequest: %s", err)
	}
	r.constructRequestHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error Scraping Reed with category=%s country=%s offset=%d: %s",
			r.category, r.country, r.offset, err,
		)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error reading response body: %s", err)
	}

	return body, nil
}

func (r *reedRequest) constructRequestHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", BASIC_PREFIX+r.apiKey)
}
