package logger

import "github.com/sirupsen/logrus"

var serverLogger = logrus.New()

// LogErr to be used to log errors
func LogErr(err error) {
	serverLogger.WithFields(logrus.Fields{
		"error": err.Error(),
	}).Fatal("Error")
}

// LogRequest to be used to logging requests
func LogRequest(method string, request interface{}, timestamp string) {
	serverLogger.WithFields(logrus.Fields{
		"method":    method,
		"request":   request,
		"timestamp": timestamp,
	}).Info("Request Information")
}

// LogInfo to log any current information
func LogInfo(timestamp string, info ...interface{}) {
	serverLogger.WithFields(logrus.Fields{
		"timestamp": timestamp,
	}).Info("Response Sent")
	serverLogger.Info(info...)
}
