package controllers_test

import (
	"errors"
	"net/http"
	"seekjob/controllers"
	"seekjob/controllers/requests"
	"seekjob/controllers/responses"
	"seekjob/models"
	"seekjob/tests/mock_cache"
	"seekjob/tests/mock_models"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestJobController(t *testing.T) {
	Convey("Job Controller", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockJobCache := mock_cache.NewMockJobHandler(mockCtrl)
		mockJobOrmer := mock_models.NewMockJobOrmer(mockCtrl)
		jobController := controllers.NewJobController(mockJobOrmer, mockJobCache)

		mockJob := &models.Job{
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
		mockDuration := time.Duration(time.Hour * 2)
		mockError := errors.New("bogus")

		expectedError := responses.NewErrorResponse(mockError.Error())

		Convey("GetJob()", func() {
			Convey("When job is cached | Should return the job value from redis", func() {
				mockJobCache.
					EXPECT().
					GetJob(mockJob.ID).
					Return(mockJob, nil)

				response, statusCode, err := jobController.GetJob(mockJob.ID)
				So(response.Job, ShouldResemble, mockJob)
				So(statusCode, ShouldEqual, http.StatusOK)
				So(err, ShouldBeNil)
			})

			Convey("When job is not cached", func() {
				Convey("When there is an error in retrieving job | Should return error", func() {
					mockJobCache.
						EXPECT().
						GetJob(mockJob.ID).
						Return(nil, nil)
					mockJobOrmer.
						EXPECT().
						Get(mockJob.ID).
						Return(nil, mockError)

					response, statusCode, err := jobController.GetJob(mockJob.ID)
					So(response, ShouldBeNil)
					So(statusCode, ShouldEqual, http.StatusInternalServerError)
					So(err, ShouldResemble, expectedError)
				})

				Convey("When there is no job with the corresponding ID | Should return nil response", func() {
					mockJobCache.
						EXPECT().
						GetJob(mockJob.ID).
						Return(nil, nil)
					mockJobOrmer.
						EXPECT().
						Get(mockJob.ID).
						Return(nil, nil)

					response, statusCode, err := jobController.GetJob(mockJob.ID)
					So(response, ShouldBeNil)
					So(statusCode, ShouldEqual, http.StatusNoContent)
					So(err, ShouldResemble, responses.NewErrorResponse("Job with the corresponding ID is not found"))
				})

				Convey("When there is a job retrieved successfully | Should return the correct response", func() {
					mockJobCache.
						EXPECT().
						GetJob(mockJob.ID).
						Return(nil, nil)
					mockJobOrmer.
						EXPECT().
						Get(mockJob.ID).
						Return(mockJob, nil)
					mockJobCache.EXPECT().SetJob(mockJob, mockDuration)

					response, statusCode, err := jobController.GetJob(mockJob.ID)
					So(response.Job, ShouldResemble, mockJob)
					So(statusCode, ShouldEqual, http.StatusOK)
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("GetJobs()", func() {
			mockParams := requests.JobRequest{
				Query:    "software",
				Category: "engineering",
				Country:  "IDN",
				PageNo:   3,
				PerPage:  50,
				Source:   "ADZ",
			}

			Convey("When there is an error in retrieveing the jobs | Should return error", func() {
				mockJobOrmer.
					EXPECT().
					GetAll("%"+mockParams.Query+"%",
						mockParams.Category,
						mockParams.Country,
						mockParams.Source,
						100, 50).
					Return(nil, mockError)

				response, statusCode, err := jobController.GetJobs(mockParams)
				So(response, ShouldBeNil)
				So(statusCode, ShouldEqual, http.StatusInternalServerError)
				So(err, ShouldResemble, expectedError)
			})

			Convey("When there are no jobs with specified parameters | Should return nil response", func() {
				mockJobOrmer.
					EXPECT().
					GetAll("%"+mockParams.Query+"%",
						mockParams.Category,
						mockParams.Country,
						mockParams.Source,
						100, 50).
					Return(nil, nil)

				response, statusCode, err := jobController.GetJobs(mockParams)
				So(response, ShouldBeNil)
				So(statusCode, ShouldEqual, http.StatusNoContent)
				So(err, ShouldResemble, responses.NewErrorResponse("Jobs not found"))
			})

			Convey("When there are jobs retrieved successfully | Should return the correct response", func() {
				mockJobs := []*models.Job{}
				for i := 0; i < 5; i++ {
					mockJobs = append(mockJobs, mockJob)
				}

				mockJobOrmer.
					EXPECT().
					GetAll("%"+mockParams.Query+"%",
						mockParams.Category,
						mockParams.Country,
						mockParams.Source,
						100, 50).
					Return(mockJobs, nil)

				response, statusCode, err := jobController.GetJobs(mockParams)
				So(response, ShouldResemble, &responses.JobsResponse{Count: len(mockJobs), Jobs: mockJobs})
				So(statusCode, ShouldEqual, http.StatusOK)
				So(err, ShouldBeNil)
			})
		})

		Convey("GetJobStatistics()", func() {
			sources := []string{"ADZ", "GIT"}
			categories := []*models.JobInfo{
				{Name: "Engineering", Total: 10},
				{Name: "Finance", Total: 5},
			}
			countries := []*models.JobInfo{
				{Name: "United States of America", Total: 15},
				{Name: "Singapore", Total: 25},
			}

			expectedResults := []*responses.JobStatistic{}
			for _, source := range sources {
				expectedResult := &responses.JobStatistic{
					Source:     source,
					Categories: categories,
					Countries:  countries,
				}
				expectedResults = append(expectedResults, expectedResult)
			}
			expectedResponse := responses.NewJobStatistics(expectedResults)

			Convey("When there is an error in retrieving sources | Should return error", func() {
				mockJobOrmer.
					EXPECT().
					GetSources().
					Return(nil, mockError)

				response, statusCode, err := jobController.GetJobsStatistics()
				So(response, ShouldBeNil)
				So(statusCode, ShouldEqual, http.StatusInternalServerError)
				So(err, ShouldResemble, expectedError)
			})

			Convey("When categories are cached", func() {
				Convey("When countries are cached | Should return the cached values", func() {
					mockJobOrmer.
						EXPECT().
						GetSources().
						Return(sources, nil)

					for _, source := range sources {
						mockJobCache.
							EXPECT().
							GetCategories(source).
							Return(categories, nil).
							AnyTimes()
						mockJobCache.
							EXPECT().
							SetCategories(source, categories, mockDuration).
							AnyTimes()
						mockJobCache.
							EXPECT().
							GetCountries(source).
							Return(countries, nil).
							AnyTimes()
						mockJobCache.
							EXPECT().
							SetCountries(source, countries, mockDuration).
							AnyTimes()
					}

					response, statusCode, err := jobController.GetJobsStatistics()
					So(response, ShouldResemble, expectedResponse)
					So(statusCode, ShouldEqual, http.StatusOK)
					So(err, ShouldBeNil)
				})

				Convey("When countries are not cached", func() {
					Convey("When there is an error in retrieving countries | Should return an error", func() {
						mockJobOrmer.
							EXPECT().
							GetSources().
							Return(sources, nil)

						for _, source := range sources {
							mockJobCache.
								EXPECT().
								GetCategories(source).
								Return(categories, nil).
								AnyTimes()
							mockJobCache.
								EXPECT().
								SetCategories(source, categories, mockDuration).
								AnyTimes()
							mockJobCache.
								EXPECT().
								GetCountries(source).
								Return(nil, nil).
								AnyTimes()
							mockJobOrmer.
								EXPECT().
								GetCountries(source).
								Return(nil, mockError).
								AnyTimes()
						}

						response, statusCode, err := jobController.GetJobsStatistics()
						So(response, ShouldBeNil)
						So(statusCode, ShouldEqual, http.StatusInternalServerError)
						So(err, ShouldNotBeNil)
					})

					Convey("When there is no error | Should return the correct response", func() {
						mockJobOrmer.
							EXPECT().
							GetSources().
							Return(sources, nil)

						for _, source := range sources {
							mockJobCache.
								EXPECT().
								GetCategories(source).
								Return(categories, nil).
								AnyTimes()
							mockJobCache.
								EXPECT().
								SetCategories(source, categories, mockDuration).
								AnyTimes()
							mockJobCache.
								EXPECT().
								GetCountries(source).
								Return(nil, nil).
								AnyTimes()
							mockJobOrmer.
								EXPECT().
								GetCountries(source).
								Return(countries, nil).
								AnyTimes()
							mockJobCache.
								EXPECT().
								SetCountries(source, countries, mockDuration).
								AnyTimes()
						}

						response, statusCode, err := jobController.GetJobsStatistics()
						So(response, ShouldNotBeNil)
						So(statusCode, ShouldEqual, http.StatusOK)
						So(err, ShouldBeNil)
					})
				})
			})

			Convey("When categories are not cached", func() {
				Convey("When there is an error in retrieving categories | Should return an error", func() {
					mockJobOrmer.
						EXPECT().
						GetSources().
						Return(sources, nil)

					for _, source := range sources {
						mockJobCache.
							EXPECT().
							GetCategories(source).
							Return(nil, nil).
							AnyTimes()
						mockJobOrmer.
							EXPECT().
							GetCategories(source).
							Return(nil, mockError).
							AnyTimes()
					}

					response, statusCode, err := jobController.GetJobsStatistics()
					So(response, ShouldBeNil)
					So(statusCode, ShouldEqual, http.StatusInternalServerError)
					So(err, ShouldResemble, expectedError)
				})

				Convey("When there is no error in retrieving categories", func() {
					Convey("When countries are cached | Should return the cached values", func() {
						mockJobOrmer.
							EXPECT().
							GetSources().
							Return(sources, nil)

						for _, source := range sources {
							mockJobCache.
								EXPECT().
								GetCategories(source).
								Return(nil, nil).
								AnyTimes()
							mockJobOrmer.
								EXPECT().
								GetCategories(source).
								Return(categories, nil).
								AnyTimes()
							mockJobCache.
								EXPECT().
								SetCategories(source, categories, mockDuration).
								AnyTimes()
							mockJobCache.
								EXPECT().
								GetCountries(source).
								Return(countries, nil).
								AnyTimes()
						}

						response, statusCode, err := jobController.GetJobsStatistics()
						So(response, ShouldResemble, expectedResponse)
						So(statusCode, ShouldEqual, http.StatusOK)
						So(err, ShouldBeNil)
					})

					Convey("When countries are not cached", func() {
						Convey("When there is an error in retrieving countries | Should return an error", func() {
							mockJobOrmer.
								EXPECT().
								GetSources().
								Return(sources, nil)

							for _, source := range sources {
								mockJobCache.
									EXPECT().
									GetCategories(source).
									Return(nil, nil).
									AnyTimes()
								mockJobOrmer.
									EXPECT().
									GetCategories(source).
									Return(categories, nil).
									AnyTimes()
								mockJobCache.
									EXPECT().
									SetCategories(source, categories, mockDuration).
									AnyTimes()
								mockJobCache.
									EXPECT().
									GetCountries(source).
									Return(nil, nil).
									AnyTimes()
								mockJobOrmer.
									EXPECT().
									GetCountries(source).
									Return(nil, mockError).
									AnyTimes()
							}

							response, statusCode, err := jobController.GetJobsStatistics()
							So(response, ShouldBeNil)
							So(statusCode, ShouldEqual, http.StatusInternalServerError)
							So(err, ShouldResemble, expectedError)
						})

						Convey("When there is no error | Should return the correct response", func() {
							mockJobOrmer.
								EXPECT().
								GetSources().
								Return(sources, nil)

							for _, source := range sources {
								mockJobCache.
									EXPECT().
									GetCategories(source).
									Return(nil, nil).
									AnyTimes()
								mockJobOrmer.
									EXPECT().
									GetCategories(source).
									Return(categories, nil).
									AnyTimes()
								mockJobCache.
									EXPECT().
									SetCategories(source, categories, mockDuration).
									AnyTimes()
								mockJobCache.
									EXPECT().
									GetCountries(source).
									Return(nil, nil).
									AnyTimes()
								mockJobOrmer.
									EXPECT().
									GetCountries(source).
									Return(countries, nil).
									AnyTimes()
								mockJobCache.
									EXPECT().
									SetCountries(source, countries, mockDuration).
									AnyTimes()
							}

							response, statusCode, err := jobController.GetJobsStatistics()
							So(response, ShouldResemble, expectedResponse)
							So(statusCode, ShouldEqual, http.StatusOK)
							So(err, ShouldBeNil)
						})
					})
				})
			})
		})
	})
}
