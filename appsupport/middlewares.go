package appsupport

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/meatballhat/negroni-logrus"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func AppMiddleware() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		reqID := GetRequestID()

		newCtx := context.WithValue(r.Context(), "reqID", reqID)
		newReq := r.WithContext(newCtx)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Add("X-Ooyala-AWS-Autoscale-Sample-App-ReqID", reqID)
		next(w, newReq)
	})
}

func AddMiddlewares(lgr *logrus.Logger, muxRouter *mux.Router) *negroni.Negroni {
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.Use(negronilogrus.NewCustomMiddleware(lgr.Level, lgr.Formatter, "awsassa"))
	n.Use(cors.New(cors.Options{}))
	n.Use(AppMiddleware())
	n.UseHandler(muxRouter)
	return n
}
