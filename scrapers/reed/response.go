package reed

type reedResponse struct {
	Results []reedResult `json:"results"`
}

type reedResult struct {
	ID          int    `json:"jobId"`
	URL         string `json:"jobUrl"`
	Title       string `json:"jobTitle"`
	Company     string `json:"employerName"`
	Description string `json:"jobDescription"`
	PostedAt    string `json:"date"`
}
