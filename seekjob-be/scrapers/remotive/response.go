package remotive

type remotiveResponse struct {
	Jobs []remotiveJob `json:"jobs"`
}

type remotiveJob struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"job_type"`
	Category    string `json:"category"`
	Company     string `json:"company"`
	PostedAt    string `json:"publication_date"`
	URL         string `json:"url"`
}
