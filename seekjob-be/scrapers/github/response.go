package github

import "time"

type githubJobsResponse struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Type        string     `json:"type"`
	Company     string     `json:"company"`
	PostedAt    *time.Time `json:"created_at"`
	URL         string     `json:"url"`
}
