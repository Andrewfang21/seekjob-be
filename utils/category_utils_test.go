package utils_test

import (
	"seekjob/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCategoryUtils(t *testing.T) {
	Convey("GetCategories()", t, func() {
		Convey("When given source exists | Should return list without error", func() {
			sources := []string{"ADZUNA", "REED", "REMOTIVE", "THEMUSE"}
			for _, source := range sources {
				result, err := utils.GetCategories(source)
				So(result, ShouldNotBeEmpty)
				So(err, ShouldBeNil)
			}
		})

		Convey("When given source exists | Should return correct list of categories", func() {
			source := "ADZUNA"
			expectedResults := []string{"accounting-finance-jobs", "engineering-jobs", "it-jobs", "pr-advertising-marketing-jobs"}
			results, err := utils.GetCategories(source)

			So(results, ShouldHaveLength, len(expectedResults))
			So(err, ShouldBeNil)

			for _, expectedResult := range expectedResults {
				So(results, ShouldContain, expectedResult)
			}
		})

		Convey("When given source does not exist | Should return error", func() {
			source := "KASKUS"
			result, err := utils.GetCategories(source)
			So(result, ShouldBeNil)
			So(err, ShouldNotBeEmpty)
		})
	})
}
