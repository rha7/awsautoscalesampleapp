package appsupport

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GetLogger() *logrus.Logger {
	lgr := logrus.New()
	lgr.SetLevel(logrus.InfoLevel)
	lgr.Out = os.Stderr
	lgr.Formatter = &logrus.JSONFormatter{
		// Base format: Mon Jan 2 15:04:05 -0700 MST 2006
		TimestampFormat: "2006-01-02 15:04:05.00000 07:00",
	}
	return lgr
}

func GetBindAddress() string {
	host := GetEnvWithDefault("AWSAUTOSCALESAMPLEAPP_HOST", "0.0.0.0")
	port := GetEnvWithDefault("AWSAUTOSCALESAMPLEAPP_PORT", "7693")
	return fmt.Sprintf("%s:%s", host, port)
}

func GetRequestID() string {
	return uuid.New().String()
}

func GetEnvWithDefault(key string, defVal string) string {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return defVal
	}
	return envVal
}
