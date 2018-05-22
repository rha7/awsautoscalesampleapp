package appsupport

import (
	"bytes"
	"os"
	"testing"

	. "github.com/rha7/awsautoscalesampleapp/goconveytesthelpers"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetLogger(t *testing.T) {
	lgr := GetLogger()
	Convey("Given a logger builder", t, func() {

		Convey("It should not be nil", func() {
			So(lgr, ShouldNotBeNil)
		})
		Convey("It should log correctly", func() {
			bfr := bytes.NewBufferString("")
			lgr.Out = bfr
			lgr.
				WithField("myfield", "myvalue").
				Info("This is Some Sample Text")
			So(bfr.String(), ShouldContainSubstring, "This is Some Sample Text")
			So(bfr.String(), ShouldContainSubstring, "myfield")
			So(bfr.String(), ShouldContainSubstring, "myvalue")
			So(bfr.String(), ShouldMatchString, "\\d{4}-\\d{2}-\\d{2}\\s+\\d{2}:\\d{2}:\\d{2}\\.\\d{5}\\s+\\d{2}:\\d{2}")
		})
	})
}

func TestGetEnvWithDefault(t *testing.T) {
	Convey("Given a GetEnv() function with defaults", t, func() {
		os.Clearenv()
		Convey("When I get an env key that exists", func() {
			os.Setenv("ENVKEY1", "Has a value")
			v := GetEnvWithDefault("ENVKEY1", "Env Key 1 Default Value")
			Convey("It should return existing value", func() {
				So(v, ShouldEqual, "Has a value")
				So(v, ShouldNotBeBlank)
			})
		})
		Convey("When I get an env key that exists, but it's empty", func() {
			os.Setenv("ENVKEY2", "")
			v := GetEnvWithDefault("ENVKEY2", "Env Key 2 Default Value")
			Convey("It should return an empty value", func() {
				So(v, ShouldEqual, "")
				So(v, ShouldBeBlank)
			})
		})
		Convey("When I get an env key that doesn't exist", func() {
			v := GetEnvWithDefault("ENVKEY3", "Env Key 3 Default Value")
			Convey("It should return a default value", func() {
				So(v, ShouldEqual, "Env Key 3 Default Value")
				So(v, ShouldNotBeBlank)
			})
		})
	})
}

func TestGetBindAddress(t *testing.T) {
	Convey("Given a Bind Address builder", t, func() {
		os.Clearenv()
		Convey("When we build an address without parameters", func() {
			bindAddress := GetBindAddress()
			Convey("We should get default values", func() {
				So(bindAddress, ShouldEqual, "0.0.0.0:7693")
			})
		})
		Convey("When we build an address with parameters", func() {
			os.Setenv("AWSAUTOSCALESAMPLEAPP_HOST", "1.2.3.4")
			os.Setenv("AWSAUTOSCALESAMPLEAPP_PORT", "4321")
			bindAddress := GetBindAddress()
			Convey("We should get specified values", func() {
				So(bindAddress, ShouldEqual, "1.2.3.4:4321")
			})
			os.Clearenv()
		})
	})
}

func TestGetRequestID(t *testing.T) {
	Convey("Given a Request ID generator", t, func() {
		Convey("When we get a new ID", func() {
			newID := GetRequestID()
			Convey("It should return a valid UUID value", func() {
				So(newID, ShouldNotBeBlank)
				So(newID, ShouldMatchString, "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
			})
		})
	})
}
