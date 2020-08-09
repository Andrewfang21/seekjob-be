package utils

import (
	"net/url"
	"strings"
)

func encodeQueryParams(str string) string {
	return strings.ReplaceAll(str, "+", "%20")
}

func ConstructUrlPath(s ...string) string {
	return strings.Join(s, "/")
}

func ConstructRequestURL(path string, params map[string]string) string {
	reqParams := url.Values{}
	for k, v := range params {
		reqParams.Set(k, v)
	}
	endpointURL, _ := url.Parse(path)
	endpointURL.RawQuery = encodeQueryParams(reqParams.Encode())

	return endpointURL.String()
}
