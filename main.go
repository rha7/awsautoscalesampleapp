package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rha7/awsautoscalesampleapp/appsupport"

	"github.com/sirupsen/logrus"
)

var (
	Version = "vDev"
	BuildTime ="Dev Time"
)

func startServer(lgr *logrus.Logger, bindAddress string, negroniRouter http.Handler) {
	var srv http.Server
	srv.Addr = bindAddress
	srv.Handler = negroniRouter

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGABRT)
		signal.Notify(sigint, syscall.SIGQUIT)
		<-sigint

		lgr.Info("shutting down server.")
		if err := srv.Shutdown(context.Background()); err != nil {
			lgr.WithError(err).Error("error occurred while shutting down server")
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		lgr.WithError(err).Error("error occurred while starting or closing listener")
	}

	<-idleConnsClosed
	lgr.Info("server shutdown.")
}

func main() {
	lgr := appsupport.GetLogger()
	bindAddress := appsupport.GetBindAddress()
	muxRoutes := appsupport.GetRoutes()
	negroniRouter := appsupport.AddMiddlewares(lgr, muxRoutes)

	lgr.
		WithField("bind_address", bindAddress).
		WithField("build_version", Version).
		WithField("build_time", BuildTime).
		Info("Starting server")
	startServer(lgr, bindAddress, negroniRouter)
}
