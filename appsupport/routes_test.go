package appsupport

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetRoutes(t *testing.T) {
	Convey("Given a router setup function", t, func() {
		router := GetRoutes()
		Convey("When I call it", func() {
			Convey("Then I should get a configured router", func() {
				So(router, ShouldNotBeNil)
			})
		})
		Convey("When I make a request to it for root path", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			router.ServeHTTP(w, r)
			Convey("Then I should get a proper response", func() {
				So(w.Code, ShouldEqual, http.StatusOK)
				So(w.Header().Get("Content-Type"), ShouldEqual, "application/json")
				So(w.Body.String(), ShouldEqual, "{\"status\":\"ok\"}")
			})
		})
	})
}
