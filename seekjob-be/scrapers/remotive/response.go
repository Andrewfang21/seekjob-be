package remotive

import "time"

type remotiveResponse struct {
	Jobs []struct {
		ID          int        `json:"id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Type        string     `json:"type"`
		Company     string     `json:"company"`
		PostedAt    *time.Time `json:"publication_date"`
		URL         string     `json:"url"`
	} `json:"jobs"`
}
