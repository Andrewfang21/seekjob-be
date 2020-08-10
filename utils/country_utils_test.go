package utils_test

import (
	"seekjob/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCountryUtils(t *testing.T) {
	Convey("GetCountries()", t, func() {
		Convey("When given source exists | Should return list without error", func() {
			sources := []string{"ADZUNA", "GITHUB", "REED"}
			for _, source := range sources {
				result, err := utils.GetCountries(source)
				So(result, ShouldNotBeEmpty)
				So(err, ShouldBeNil)
			}
		})

		Convey("When given source exists | Should return correct list of countries", func() {
			source := "REED"
			expectedResults := []string{"America", "Australia", "Canada", "India", "Singapore"}
			results, err := utils.GetCountries(source)

			So(results, ShouldHaveLength, len(expectedResults))
			So(err, ShouldBeNil)

			for _, expectedResult := range expectedResults {
				So(results, ShouldContain, expectedResult)
			}
		})

		Convey("When given source does not exist | Should return error", func() {
			source := "KASKUS"
			result, err := utils.GetCountries(source)
			So(result, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})
	})

	Convey("GetCountry()", t, func() {
		Convey("When given source exists | Should return the correct country", func() {
			sources := []string{"USA", "IDN", "SGP", "CAN"}
			expectedResults := []string{"United States of America", "Indonesia", "Singapore", "Canada"}
			for i := 0; i < len(sources); i++ {
				result, err := utils.GetCountry(sources[i])
				So(result, ShouldEqual, expectedResults[i])
				So(err, ShouldBeNil)
			}
		})
		Convey("When given source does not exist | Should return error", func() {
			source := "HKG"
			result, err := utils.GetCountry(source)
			So(result, ShouldBeBlank)
			So(err, ShouldNotBeNil)
		})
	})

	Convey("GetCountryCode()", t, func() {
		Convey("When given country exists | Should return the correct country ISO Code", func() {
			countries := []string{"America", "United States of America", "Singapore", "Remote", "Canada"}
			expectedResults := []string{"USA", "USA", "SGP", "REM", "CAN"}
			for i := 0; i < len(countries); i++ {
				result, err := utils.GetCountryCode(countries[i])
				So(result, ShouldEqual, expectedResults[i])
				So(err, ShouldBeNil)
			}
		})
		Convey("When given country does not exist | Should return error", func() {
			countries := []string{"Hong Kong", "Taiwan", "France", "Spain"}
			for _, country := range countries {
				result, err := utils.GetCountryCode(country)
				So(result, ShouldBeBlank)
				So(err, ShouldNotBeNil)
			}
		})
	})
}
