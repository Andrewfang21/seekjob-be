package utils_test

import (
	"seekjob/utils"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStringUtils(t *testing.T) {
	Convey("encodeQueryParams()", t, func() {
		str := "software+engineering"
		expectedResult := "software%20engineering"
		So(utils.EncodeQueryParams(str), ShouldEqual, expectedResult)
	})

	Convey("ConstructURLPath()", t, func() {
		s := []string{"a", "b", "c", "d"}
		expectedResult := strings.Join(s, "/")
		So(utils.ConstructURLPath(s...), ShouldEqual, expectedResult)
	})

	Convey("ConstructRequestURL()", t, func() {
		path := "www.google.com/search"
		params := map[string]string{
			"keyA": "val1 val2",
			"keyB": "val3+val4",
			"keyC": "val5",
		}
		expectedResult := path + "?keyA=val1%20val2&keyB=val3%2Bval4&keyC=val5"
		So(utils.ConstructRequestURL(path, params), ShouldEqual, expectedResult)
	})
}
