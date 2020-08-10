package models_test

import (
	"seekjob/models"
	"strconv"
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

		insertQueryCmd := func(item *models.Job) error {
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
				) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			`
			_, err := ormer.Exec(queryString,
				item.ID,
				item.URL,
				item.Title,
				item.Company,
				item.Description,
				item.Category,
				item.Country,
				item.Type,
				item.PostedAt,
				item.Source,
			)
			return err
		}

		Convey("Get()", func() {
			Convey("Should return the correct row", func() {
				err := insertQueryCmd(mockJobItem)
				So(err, ShouldBeNil)

				result, err := j.Get(mockJobItem.ID)
				So(err, ShouldBeNil)
				So(result, ShouldResemble, mockJobItem)
			})

			Convey("When multiple rows exist | Should return the correct row", func() {
				otherMockJobItem := &models.Job{
					ID:          "2",
					URL:         "www.facebook.com",
					Title:       "Data Engineering Intern",
					Category:    "Engineering",
					Company:     "Facebook",
					Country:     "UK",
					Description: "Another Internship",
					PostedAt:    time.Now().Unix(),
					Source:      "ADZ",
					Type:        "Entry Level",
				}

				err := insertQueryCmd(mockJobItem)
				So(err, ShouldBeNil)
				err = insertQueryCmd(otherMockJobItem)
				So(err, ShouldBeNil)

				result, err := j.Get(mockJobItem.ID)
				So(result, ShouldResemble, mockJobItem)

				result, err = j.Get(otherMockJobItem.ID)
				So(result, ShouldResemble, otherMockJobItem)
			})

			Convey("Should return nil when there are no data", func() {
				result, err := j.Get(mockJobItem.ID)
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
			})
		})

		Convey("GetAll()", func() {
			mockItems := make([]models.Job, 20)
			for i := 0; i < 20; i++ {
				mockItems[i] = *mockJobItem
				mockItems[i].ID = strconv.Itoa(i)
			}
			insertMockItems := func() {
				for _, item := range mockItems {
					err := insertQueryCmd(&item)
					So(err, ShouldBeNil)
				}
			}

			Convey("Should return the correct number of rows", func() {
				insertMockItems()

				query := "%"
				category := mockJobItem.Category
				country := mockJobItem.Country
				source := mockJobItem.Source
				offset := 0
				limit := 10

				results, err := j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldHaveLength, limit)
			})

			Convey("Should return the correct number of rows and offset", func() {
				insertMockItems()
				query := "%"
				category := mockJobItem.Category
				country := mockJobItem.Country
				source := mockJobItem.Source
				offset := 3
				limit := 10

				results, err := j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldHaveLength, limit)

				for i := 0; i < len(results); i++ {
					So(results[i], ShouldResemble, &mockItems[i+offset])
				}
			})

			Convey("When given the description substring | Should return the correct rows", func() {
				insertMockItems()

				category := mockJobItem.Category
				country := mockJobItem.Country
				source := mockJobItem.Source
				offset := 0
				limit := 20

				query := "%ntern%"
				results, err := j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldHaveLength, limit)

				query = "inte%"
				results, err = j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldHaveLength, limit)

				query = "%SHIP"
				results, err = j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldHaveLength, limit)

				query = "qwertyuiop"
				results, err = j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldBeEmpty)
			})

			Convey("When given country | Should return the correct rows", func() {
				insertMockItems()

				query := "%"
				category := mockJobItem.Category
				source := mockJobItem.Source
				offset := 0
				limit := 10

				country := "usa"
				results, err := j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldHaveLength, limit)

				country = "uSA"
				results, err = j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldHaveLength, limit)

				country = "IDN"
				results, err = j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldHaveLength, 0)
			})

			Convey("When given category | Should return the correct rows", func() {
				insertMockItems()

				query := "%"
				country := mockJobItem.Country
				source := mockJobItem.Source
				offset := 0
				limit := 10

				category := "ENGINEERING"
				results, err := j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldHaveLength, limit)

				category = "engineering"
				results, err = j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldHaveLength, limit)

				category = "Marketing"
				results, err = j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldHaveLength, 0)
			})

			Convey("When given source | Should return the correct rows", func() {
				insertMockItems()

				query := "%"
				country := mockJobItem.Country
				category := mockJobItem.Category
				offset := 0
				limit := 10

				source := "GIT"
				results, err := j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldHaveLength, limit)

				source = "git"
				results, err = j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldHaveLength, limit)

				source = "KASKUS"
				results, err = j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldHaveLength, 0)
			})

			Convey("Should return empty slice when there are no data", func() {
				query := "%"
				category := "%"
				country := "%"
				source := "%"
				offset := 0
				limit := 1000000

				results, err := j.GetAll(query, category, country, source, offset, limit)
				So(err, ShouldBeNil)
				So(results, ShouldBeEmpty)
			})
		})

		Convey("Upsert()", func() {
			otherMockItem := &models.Job{
				ID:          mockJobItem.ID,
				URL:         "www.twitter.com",
				Title:       "Data Science Intern",
				Company:     "Twitter",
				Description: "Internship at Twitter",
				Category:    "Data",
				Country:     "SGP",
				PostedAt:    time.Now().Unix(),
				Type:        "Entry Level",
				Source:      "GIT",
			}

			scanQueryResult := func(id string) []*models.Job {
				queryString := `
					SELECT * FROM jobs
					WHERE id=$1
				`
				queryResult, err := ormer.Query(queryString, id)
				So(err, ShouldBeNil)
				So(queryResult, ShouldNotBeNil)

				defer queryResult.Close()

				var jobs []*models.Job
				for queryResult.Next() {
					var job models.Job
					err = queryResult.Scan(
						&job.ID,
						&job.URL,
						&job.Title,
						&job.Company,
						&job.Description,
						&job.Category,
						&job.Country,
						&job.Type,
						&job.PostedAt,
						&job.Source,
					)
					So(err, ShouldBeNil)
					jobs = append(jobs, &job)
				}
				return jobs
			}

			Convey("Should insert the correct row", func() {
				err := j.Upsert(*mockJobItem)
				So(err, ShouldBeNil)

				results := scanQueryResult(mockJobItem.ID)
				So(results, ShouldHaveLength, 1)
				So(results[0], ShouldResemble, mockJobItem)
			})

			Convey("Should update on ID conflict", func() {
				err := j.Upsert(*mockJobItem)
				So(err, ShouldBeNil)

				err = j.Upsert(*otherMockItem)
				So(err, ShouldBeNil)

				results := scanQueryResult(mockJobItem.ID)
				So(results, ShouldHaveLength, 1)
				So(results[0], ShouldResemble, otherMockItem)
			})

			Convey("Should insert multiple distinct rows withour error", func() {
				err := j.Upsert(*mockJobItem)
				So(err, ShouldBeNil)

				otherMockItem.ID = "2"
				err = j.Upsert(*otherMockItem)
				So(err, ShouldBeNil)

				queryString := `SELECT * FROM jobs`
				queryResult, err := ormer.Query(queryString)
				So(err, ShouldBeNil)

				defer queryResult.Close()

				var jobs []*models.Job
				for queryResult.Next() {
					var job models.Job
					err = queryResult.Scan(
						&job.ID,
						&job.URL,
						&job.Title,
						&job.Company,
						&job.Description,
						&job.Category,
						&job.Country,
						&job.Type,
						&job.PostedAt,
						&job.Source,
					)
					So(err, ShouldBeNil)
					jobs = append(jobs, &job)
				}

				So(jobs, ShouldHaveLength, 2)
				So(jobs[0], ShouldResemble, mockJobItem)
				So(jobs[1], ShouldResemble, otherMockItem)
			})
		})
	})
}
