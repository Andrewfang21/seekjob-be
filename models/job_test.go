package models_test

import (
	"seekjob/models"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJob(t *testing.T) {
	Convey("job", t, func() {
		ormer := setUpDatabase(t, "jobs")
		j := models.NewJobOrmer(ormer)

		mockJobItem := &models.Job{
			ID:          "1",
			URL:         "www.google.com",
			Title:       "Software Engineering Intern",
			Category:    "Engineering",
			Company:     "Google",
			Country:     "USA",
			Description: "Internship",
			PostedAt:    time.Now().Unix(),
			Source:      "GIT",
			Type:        "Internship",
		}

		Convey("Get()", func() {
			Convey("Should query the correct row", func() {
				queryString := `
					INSERT INTO jobs(
						id, 
						url, 
						title, 
						company, 
						description, 
						category, 
						country, 
						type, 
						time, 
						source
					)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
				`
				_, err := ormer.Exec(queryString,
					mockJobItem.ID,
					mockJobItem.URL,
					mockJobItem.Title,
					mockJobItem.Company,
					mockJobItem.Description,
					mockJobItem.Category,
					mockJobItem.Country,
					mockJobItem.Type,
					mockJobItem.PostedAt,
					mockJobItem.Source,
				)
				So(err, ShouldBeNil)

				result, err := j.Get(mockJobItem.ID)
				So(err, ShouldBeNil)
				So(result, ShouldResemble, mockJobItem)
			})
		})
	})
}
