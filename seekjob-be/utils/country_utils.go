package utils

var GITHUB_JOBS_COUNTRIES = []string{
	"Australia",
	"Canada",
	"Indonesia",
	"India",
	"Singapore", // As of now, there are no jobs based in Singapore that are listed in GithubJobs
	"United States of America",
}

var REED_COUNTRIES = []string{
	"America",
	"Australia",
	"Canada",
	"India",
	"Singapore",
}

var ADZUNA_COUNTRIES_CODE_MAP = map[string]string{
	"sg": "Singapore",
	"us": "United States of America",
	"au": "Australia",
	"in": "India",
}

var COUNTRIES_STRING_MAP = map[string]string{
	"Singapore":                "SGP",
	"Indonesia":                "IDN",
	"India":                    "IND",
	"United States of America": "USA",
	"America":                  "USA", // To handle Reed API, since Reed API does not support location `United State of America`
	"Canada":                   "CAN",
	"Australia":                "AUS",
	"Remote":                   "REM",
}

var COUNTRIES_ISO_CODE_MAP = map[string]string{
	"SGP": "Singapore",
	"IDN": "Indonesia",
	"IND": "India",
	"USA": "United States of America",
	"CAN": "Canada",
	"AUS": "Australia",
	"REM": "Remote",
}
