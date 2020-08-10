package utils

import (
	"net/url"
	"strings"
)

// EncodeQueryParams export private function `encodeQueryParams` for testing purpose
var EncodeQueryParams = encodeQueryParams

func encodeQueryParams(str string) string {
	return strings.ReplaceAll(str, "+", "%20")
}

// ConstructURLPath joins all given strings with `/`
// Ex: www.google.com`/`search`/`1
func ConstructURLPath(s ...string) string {
	return strings.Join(s, "/")
}

// ConstructRequestURL encodes given params and append it with the given path
func ConstructRequestURL(path string, params map[string]string) string {
	reqParams := url.Values{}
	for k, v := range params {
		reqParams.Set(k, v)
	}
	endpointURL, _ := url.Parse(path)
	endpointURL.RawQuery = encodeQueryParams(reqParams.Encode())

	return endpointURL.String()
}
