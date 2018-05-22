package appsupport

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/rha7/awsautoscalesampleapp/goconveytesthelpers"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func TestAddMiddlewares(t *testing.T) {
	Convey("Given a mux router, a logrus logger, and a middleware manager", t, func() {
		b := bytes.NewBufferString("")
		r := mux.NewRouter()
		l := logrus.New()
		l.Out = b
		Convey("When I call it to decorate a router", func() {
			n := AddMiddlewares(l, r)
			Convey("It should be valid and decorated", func() {
				So(n, ShouldNotBeNil)
				So(n.Handlers(), ShouldHaveLength, 6) // Currently 5 middlewares plus router
			})
		})
	})
}

func TestAppMiddleware(t *testing.T) {
	Convey("Given an app middleware builder, and a test request", t, func() {
		tr, err := http.NewRequest("GET", "", nil)
		if err != nil {
			panic(err)
		}
		Convey("When we create a new middleware for app", func() {
			mw := AppMiddleware()
			tw := &httptest.ResponseRecorder{}
			contentTypeFromHdr := "none"
			requestIDFromHdr := "none"
			requestIDFromCtx := "none"
			mw(tw, tr, func(w http.ResponseWriter, r *http.Request) {
				var ok bool
				ctx := r.Context()
				reqIDifc := ctx.Value("reqID")
				requestIDFromCtx, ok = reqIDifc.(string)
				if !ok {
					requestIDFromCtx = "error"
				}
				requestIDFromHdr = tw.Header().Get("X-Ooyala-AWS-Autoscale-Sample-App-ReqID")
				contentTypeFromHdr = tw.Header().Get("Content-Type")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("done!"))
			})
			Convey("It should add ReqID and Content type headers", func() {
				So(requestIDFromHdr, ShouldNotBeBlank)
				So(requestIDFromHdr, ShouldMatchString, "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
				So(contentTypeFromHdr, ShouldNotBeBlank)
				So(contentTypeFromHdr, ShouldEqual, "application/json")
			})
			Convey("It should add a the request ID to the request context", func() {
				So(requestIDFromCtx, ShouldNotBeBlank)
				So(requestIDFromCtx, ShouldMatchString, "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
			})
			Convey("It should have request ID from context and from header be the same", func() {
				So(requestIDFromCtx, ShouldEqual, requestIDFromHdr)
			})
		})
	})
}
