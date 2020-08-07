package themuse

import "time"

type theMuseResponse struct {
	Results []theMuseResult `json:"results"`
}

type theMuseResult struct {
	ID  int `json:"id"`
	URL struct {
		LandingPage string `json:"landing_page"`
	} `json:"refs"`
	Title   string `json:"name"`
	Company struct {
		Name string `json:"name"`
	} `json:"company"`
	Description string     `json:"contents"`
	PostedAt    *time.Time `json:"publication_date"`
}
