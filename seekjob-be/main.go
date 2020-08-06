package main

import (
	"fmt"
	"log"
	"os"
	"seekjob/cache"
	"seekjob/config"
	"seekjob/controllers"
	"seekjob/database"
	"seekjob/models"
	"seekjob/redis"
	"seekjob/scrapers"
	"seekjob/scrapers/adzuna"
	"seekjob/scrapers/github"
	"seekjob/scrapers/remotive"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	r := gin.New()
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	redisHandler := redis.GetHandler()
	postgresHandler := database.GetHandler()
	jobCacheHandler := cache.NewJobCacheHandler(redisHandler)
	jobOrmer := models.NewJobOrmer(postgresHandler)
	jobController := controllers.NewJobController(jobOrmer, jobCacheHandler)

	scraperCfg := config.ScraperCfg
	adzunaScraper := adzuna.NewAdzunaScraperHandler(jobOrmer, scraperCfg.Adzuna)
	githubJobsScraper := github.NewGithubJobsScraperHandler(jobOrmer)
	remotiveScraper := remotive.NewRemotiveScraperHandler(jobOrmer)
	scraper := scrapers.NewScraperHandler(adzunaScraper, githubJobsScraper, remotiveScraper)

	c := cron.New()
	c.AddFunc("@every 0h0m2s", scraper.ScrapeJobs)
	c.Start()
	// scraper.ScrapeJobs()

	r.GET("/api/jobs", jobController.GetJobs)
	r.GET("/api/jobs/id/:id", jobController.GetJob)
	r.GET("/api/job/stats", jobController.GetJobsStatistics)
	r.GET("/api/jobs/category/:category", jobController.GetJobsByCategory)
	r.GET("/api/jobs/location/:location", jobController.GetJobsByLocation)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := r.Run(port); err != nil {
		log.Fatalf("[ERROR] Fatal error running port %s: %s", port, err)
	}
}
