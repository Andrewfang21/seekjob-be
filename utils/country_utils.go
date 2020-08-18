package utils

import (
	"fmt"
)

var adzunaCountries = []string{
	"sg",
	"us",
	"au",
	"in",
}

var githubJobsCountries = []string{
	"Australia",
	"Canada",
	"Indonesia",
	"India",
	"Singapore", // As of now, there are no jobs based in Singapore that are listed in GithubJobs
	"United States of America",
}

var reedCountries = []string{
	"America",
	"Australia",
	"Canada",
	"India",
	"Singapore",
}

// GetCountries return list of string according to the source
func GetCountries(source string) ([]string, error) {
	list := []string{}
	if source == "ADZUNA" {
		list = append(list, adzunaCountries...)
	} else if source == "GITHUB" {
		list = append(list, githubJobsCountries...)
	} else if source == "REED" {
		list = append(list, reedCountries...)
	} else {
		return nil, fmt.Errorf("Source %s does not exist", source)
	}
	return list, nil
}

var countryMap = map[string]string{
	"Singapore":                "SGP",
	"Indonesia":                "IDN",
	"India":                    "IND",
	"United States of America": "USA",
	"America":                  "USA", // To handle Reed API, since Reed API does not support location `United State of America`
	"Canada":                   "CAN",
	"Australia":                "AUS",
	"Remote":                   "REM",
	"Flexbile / Remote":        "REM",
}

var countryISOCodeMap = map[string]string{
	"SGP": "Singapore",
	"IDN": "Indonesia",
	"IND": "India",
	"US":  "United States of America", // To handle Adzuna API location format
	"USA": "United States of America",
	"CAN": "Canada",
	"AUS": "Australia",
	"REM": "Remote",
}

// GetCountry returns country according to the given country iso code
func GetCountry(code string) (string, error) {
	val, ok := countryISOCodeMap[code]
	if !ok {
		return "", fmt.Errorf("Country with code %s does not exist", code)
	}
	return val, nil
}

// GetCountryCode returns country iso code according to the given country
func GetCountryCode(country string) (string, error) {
	val, ok := countryMap[country]
	if !ok {
		return "", fmt.Errorf("Country %s does not exist", country)
	}
	return val, nil
}
