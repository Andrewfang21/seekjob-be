package remotive

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"seekjob/utils"
)

const API_BASE_URL = "https://remotive.io/api"

type remotiveRequestable interface {
	callEndpoint(method string) ([]byte, error)
	constructRequestHeaders(req *http.Request)
}

type remotiveRequest struct {
	category string
}

func newRemotiveRequest(category string) remotiveRequestable {
	return &remotiveRequest{category: category}
}

/*
	API Docs: https://remotive.io/api-documentation
	Params: @required category

	By default, Remotive does not support pagination
*/
func (r *remotiveRequest) callEndpoint(method string) ([]byte, error) {
	path := utils.ConstructUrlPath(API_BASE_URL, "remote-jobs")
	params := map[string]string{"category": r.category}

	url := utils.ConstructRequestURL(path, params)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error creating NewRequest: %s", err)
	}
	r.constructRequestHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error scraping Remotive with category=%s: %s",
			r.category, err,
		)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error reading response body: %s", err)
	}

	return body, nil
}

func (r *remotiveRequest) constructRequestHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}
