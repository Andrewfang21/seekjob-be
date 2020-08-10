package responses

// JobStatistic model
type JobStatistic struct {
	Results []struct {
		Source     string `json:"source"`
		Categories []struct {
			Name  string `json:"name"`
			Total int    `json:"job_count"`
		} `json:"categories"`
		Countries []struct {
			Name  string `json:"name"`
			Total int    `json:"job_count"`
		} `json:"countries"`
	} `json:"results"`
}
