package github

type githubJobsResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Company     string `json:"company"`
	Country     string
	PostedAt    string `json:"created_at"`
	URL         string `json:"url"`
}
