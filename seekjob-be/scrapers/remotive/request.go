package remotive

import "fmt"

const API_BASE_URL string = "https://remotive.io/api"

type remotiveRequestable interface {
	constructEndpoints() string
}

type remotiveRequest struct {
	tag string
}

func NewRemotiveRequest(tag string) remotiveRequestable {
	return &remotiveRequest{tag: tag}
}

func (r *remotiveRequest) constructEndpoints() string {
	apiEndpoint := fmt.Sprintf("%s/remote-jobs?category=%s",
		API_BASE_URL,
		r.tag,
	)
	return apiEndpoint
}
