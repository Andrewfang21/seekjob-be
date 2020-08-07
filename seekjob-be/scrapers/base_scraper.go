package scrapers

import (
	"runtime"
	"seekjob/scrapers/adzuna"
	"seekjob/scrapers/github"
	"seekjob/scrapers/reed"
	"seekjob/scrapers/remotive"
)

type scraper struct {
	adzunaScraper    adzuna.Handler
	githubJobScraper github.Handler
	reedScraper      reed.Handler
	remotiveScraper  remotive.Handler
}

// Scraper defines the operations of scraper
type Scraper interface {
	ScrapeJobs()
}

// NewScraperHandler returns handler for all jobs API
func NewScraperHandler(
	adzunaScraper adzuna.Handler,
	githubJobsScaper github.Handler,
	reedScraper reed.Handler,
	remotiveScraper remotive.Handler) Scraper {
	return &scraper{
		adzunaScraper:    adzunaScraper,
		githubJobScraper: githubJobsScaper,
		reedScraper:      reedScraper,
		remotiveScraper:  remotiveScraper,
	}
}

func (s *scraper) ScrapeJobs() {
	runtime.GOMAXPROCS(3)

	// go s.adzunaScraper.ScrapeJobs()
	// go s.githubJobScraper.ScrapeJobs()
	// go s.remotiveScraper.ScrapeJobs()
	// go s.reedScraper.ScrapeJobs()
}
