package cache_test

import (
	"errors"
	"seekjob/cache"
	"seekjob/models"
	"seekjob/tests/mock_redis"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestJobCache(t *testing.T) {
	Convey("Job Cache", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		redisHandler := mock_redis.NewMockHandler(mockCtrl)
		handler := cache.NewJobCacheHandler(redisHandler)

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
		mockCountries := []*models.JobInfo{
			{
				Name:  "USA",
				Total: 50,
			},
			{
				Name:  "SGP",
				Total: 30,
			},
		}
		mockCategories := []*models.JobInfo{
			{
				Name:  "Engineering",
				Total: 100,
			},
			{
				Name:  "Accounting",
				Total: 20,
			},
		}
		mockError := errors.New("bogus")

		jobsRedisKey := func(ID string) string {
			return "jobs:id:" + ID
		}
		countriesRedisKey := func(source string) string {
			return "countries:source:" + source
		}
		categoriesRedisKey := func(source string) string {
			return "categories:source:" + source
		}

		Convey("GetJob()", func() {
			Convey("When there is an error in getting value | Should return error", func() {
				redisHandler.
					EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(false, mockError)

				result, err := handler.GetJob(gomock.Any().String())
				So(err, ShouldEqual, mockError)
				So(result, ShouldBeNil)
			})

			Convey("When key does not exist | Should return nil", func() {
				redisHandler.
					EXPECT().
					Get(jobsRedisKey(mockJob.ID), gomock.Any()).
					Return(false, nil)

				result, err := handler.GetJob(mockJob.ID)
				So(err, ShouldBeNil)
				So(result, ShouldBeNil)
			})

			Convey("When key does exist | Should return the correct value", func() {
				redisHandler.
					EXPECT().
					Get(jobsRedisKey(mockJob.ID), gomock.Any()).
					SetArg(1, *mockJob).
					Return(true, nil)

				result, err := handler.GetJob(mockJob.ID)
				So(err, ShouldBeNil)
				So(result, ShouldResemble, mockJob)
			})
		})

		Convey("GetCountries()", func() {
			Convey("When there is an error in getting value | Should return error", func() {
				redisHandler.
					EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(false, mockError)

				result, err := handler.GetCountries(gomock.Any().String())
				So(err, ShouldEqual, mockError)
				So(result, ShouldBeNil)
			})

			Convey("When key does not exist | Should return nil", func() {
				redisHandler.
					EXPECT().
					Get(countriesRedisKey(gomock.Any().String()), gomock.Any()).
					Return(false, nil)

				results, err := handler.GetCountries(gomock.Any().String())
				So(err, ShouldBeNil)
				So(results, ShouldBeNil)
			})

			Convey("When key does exist | Should return the correct value", func() {
				redisHandler.
					EXPECT().
					Get(countriesRedisKey("GIT"), gomock.Any()).
					SetArg(1, mockCountries).
					Return(true, nil)

				results, err := handler.GetCountries("GIT")
				So(err, ShouldBeNil)
				So(results, ShouldResemble, mockCountries)
			})
		})

		Convey("GetCategories()", func() {
			Convey("When there is an error in getting value | Should return error", func() {
				redisHandler.
					EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(false, mockError)

				result, err := handler.GetCountries(gomock.Any().String())
				So(err, ShouldEqual, mockError)
				So(result, ShouldBeNil)
			})

			Convey("When key does not exist | Should return nil", func() {
				redisHandler.
					EXPECT().
					Get(categoriesRedisKey(gomock.Any().String()), gomock.Any()).
					Return(false, nil)

				results, err := handler.GetCategories(gomock.Any().String())
				So(err, ShouldBeNil)
				So(results, ShouldBeNil)
			})

			Convey("When key does exist | Should return the correct value", func() {
				redisHandler.
					EXPECT().
					Get(categoriesRedisKey("GIT"), gomock.Any()).
					SetArg(1, mockCategories).
					Return(true, nil)

				results, err := handler.GetCategories("GIT")
				So(err, ShouldBeNil)
				So(results, ShouldResemble, mockCategories)
			})
		})

		Convey("SetJob()", func() {
			Convey("When there is an error in storing value | Should return error", func() {

			})

			Convey("When there is no error in storing value | Should return nil", func() {

			})
		})

		Convey("SetCountries()", func() {

		})

		Convey("SetCategories()", func() {

		})
	})
}
