package themuse

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"seekjob/utils"
	"strconv"
)

const API_BASE_URL string = "https://www.themuse.com/api/public/jobs"

type themuseRequestable interface {
	callEndpoint(method string) ([]byte, error)
	constructRequestHeaders(req *http.Request)
}

type themuseRequest struct {
	apiKey      string
	category    string
	currentPage int
}

func newTheMuseRequest(
	apiKey, category string,
	currentPage int) themuseRequestable {
	return &themuseRequest{
		apiKey:      apiKey,
		category:    category,
		currentPage: currentPage,
	}
}

/*
	API Docs: https://www.themuse.com/developers/api/v2?ref=apilist.fun
	Params:	@required api_key
			@required page
			@optional category

	By default, TheMuse API returns at most 20 results per page
*/
func (t *themuseRequest) callEndpoint(method string) ([]byte, error) {
	currentPage := strconv.Itoa(t.currentPage)
	path := API_BASE_URL
	params := map[string]string{
		"api_key":  t.apiKey,
		"category": t.category,
		"page":     currentPage,
	}

	url := utils.ConstructRequestURL(path, params)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error creating NewRequest: %s", err)
	}
	t.constructRequestHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error Scraping TheMuse with category=%s page=%s: %s",
			t.category, currentPage, err,
		)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error reading response body: %s", err)
	}

	return body, nil
}

func (t *themuseRequest) constructRequestHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")

}
