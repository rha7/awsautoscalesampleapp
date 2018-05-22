package goconveytesthelpers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShouldMatchString(t *testing.T) {
	Convey("Given a Goconvey regular expression matcher", t, func() {
		Convey("When I compare a matching string", func() {
			retVal := ShouldMatchString("Alfa Value", "Alfa\\s+Value")
			Convey("Then I should get a blank return value", func() {
				So(retVal, ShouldBeBlank)
			})
		})
		Convey("When I compare a non-matching string", func() {
			retVal := ShouldMatchString("Alfa Value", "Beta\\s+Value")
			Convey("Then I should get an error message", func() {
				So(retVal, ShouldNotBeBlank)
				So(retVal, ShouldContainSubstring, "does not match expression")
			})
		})
		Convey("When I compare a non-string value", func() {
			retVal := ShouldMatchString(true, "Beta\\s+Value")
			Convey("Then I should get an error message", func() {
				So(retVal, ShouldNotBeBlank)
				So(retVal, ShouldContainSubstring, "invalid actual")
				So(retVal, ShouldContainSubstring, "not a string")
			})
		})
		Convey("When I compare a string value against multiple regular expressions", func() {
			retVal := ShouldMatchString("Alfa Value", "Beta\\s+Value", "Gamma\\s+Value")
			Convey("Then I should get an error message", func() {
				So(retVal, ShouldNotBeBlank)
				So(retVal, ShouldContainSubstring, "missing or more than one")
			})
		})
		Convey("When I compare a string value against a non-string value", func() {
			retVal := ShouldMatchString("Alfa Value", 5)
			Convey("Then I should get an error message", func() {
				So(retVal, ShouldNotBeBlank)
				So(retVal, ShouldContainSubstring, "invalid expected")
				So(retVal, ShouldContainSubstring, "not a string")
			})
		})
		Convey("When I compare a string value against an invalid regular expression", func() {
			retVal := ShouldMatchString("Alfa Value", "This)is not[a good) regexp")
			Convey("Then I should get an error message", func() {
				So(retVal, ShouldNotBeBlank)
				So(retVal, ShouldContainSubstring, "invalid expected")
				So(retVal, ShouldContainSubstring, "invalid regular expression")
			})
		})
	})
}
