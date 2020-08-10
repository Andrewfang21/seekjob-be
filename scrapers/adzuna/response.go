package adzuna

import "time"

type adzunaResponse struct {
	Results []adzunaResult `json:"results"`
}

type adzunaResult struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Descripton string `json:"description"`
	Type       string `json:"type"`
	Company    struct {
		Name string `json:"display_name"`
	} `json:"company"`
	Category struct {
		Label string `json:"label"`
	} `json:"category"`
	Country struct {
		List []string `json:"area"`
	} `json:"location"`
	PostedAt *time.Time `json:"created"`
	URL      string     `json:"redirect_url"`
}
