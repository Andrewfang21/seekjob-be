package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"seekjob/cache"
	"seekjob/config"
	"seekjob/controllers"
	"seekjob/controllers/requests"
	"seekjob/controllers/responses"
	"seekjob/database"
	"seekjob/models"
	"seekjob/redis"
	"seekjob/scrapers"
	"seekjob/scrapers/adzuna"
	"seekjob/scrapers/github"
	"seekjob/scrapers/reed"
	"seekjob/scrapers/remotive"
	"seekjob/scrapers/themuse"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("pas di bawah main")
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
	reedScraper := reed.NewReedScraperHandler(jobOrmer, scraperCfg.Reed)
	remotiveScraper := remotive.NewRemotiveScraperHandler(jobOrmer)
	theMuseScraper := themuse.NewTheMuseScraperHandler(jobOrmer, scraperCfg.TheMuse)
	scraper := scrapers.NewScraperHandler(
		adzunaScraper,
		githubJobsScraper,
		reedScraper,
		remotiveScraper,
		theMuseScraper,
	)

	// c := cron.New()
	// c.AddFunc("@every 0h0m2s", scraper.ScrapeJobs)
	// c.Start()
	scraper.ScrapeJobs()
	fmt.Println("before router")

	configureRouters(r, jobController)

	fmt.Println("after router")

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := r.Run(port); err != nil {
		log.Fatalf("[ERROR] Fatal error running port %s: %s", port, err)
	}
}

func configureRouters(r *gin.Engine, j controllers.JobController) {
	r.GET("/api/jobs/id/:id", func(c *gin.Context) {
		id := c.Param("id")
		response, statusCode, err := j.GetJob(id)
		if err != nil {
			c.JSON(statusCode, err)
			return
		}
		c.JSON(statusCode, response)
	})

	r.GET("/api/jobs", func(c *gin.Context) {
		var params requests.JobRequest
		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, responses.NewErrorResponse(err.Error()))
			return
		}

		response, statusCode, err := j.GetJobs(params)
		if err != nil {
			c.JSON(statusCode, err)
			return
		}
		c.JSON(statusCode, response)
	})

	r.GET("/api/jobs/stats", func(c *gin.Context) {
		response, statusCode, err := j.GetJobsStatistics()
		if err != nil {
			c.JSON(statusCode, err)
			return
		}
		c.JSON(statusCode, response)
	})
}
