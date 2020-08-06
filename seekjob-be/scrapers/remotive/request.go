package remotive

import (
	"seekjob/utils"
)

const API_BASE_URL = "https://remotive.io/api"

type remotiveRequestable interface {
	constructEndpoints() string
}

type remotiveRequest struct {
	category string
}

func newRemotiveRequest(category string) remotiveRequestable {
	return &remotiveRequest{category: category}
}

/*
	Params: @required category
*/
func (r *remotiveRequest) constructEndpoints() string {
	endpoint :=
		utils.ConstructAPIPath(API_BASE_URL, "remote-jobs") +
			"?" +
			utils.ConstructAPIQuery("category", r.category)
	return endpoint
}
