package scrapers

import (
	"fmt"
	"seekjob/scrapers/adzuna"
	"seekjob/scrapers/github"
	"seekjob/scrapers/remotive"
)

type scraper struct {
	adzunaScraper    adzuna.Handler
	githubJobScraper github.Handler
	remotiveScraper  remotive.Handler
}

type scraperable interface {
	ScrapeJobs()
}

func NewScraperHandler(
	adzunaScraper adzuna.Handler,
	githubJobsScaper github.Handler,
	remotiveScraper remotive.Handler,
) scraperable {
	return &scraper{
		adzunaScraper:    adzunaScraper,
		githubJobScraper: githubJobsScaper,
		remotiveScraper:  remotiveScraper,
	}
}

func (s *scraper) ScrapeJobs() {
	fmt.Println("Inside ScrapeJobs() in base_scraper.go")
}
