package logger

import "github.com/sirupsen/logrus"

var serverLogger = logrus.New()

// LogErr to be used to log errors
func LogErr(err error, errMsg string) {
	serverLogger.WithFields(logrus.Fields{
		"error": err.Error(),
		"message": errMsg,
	}).Fatal("There has been an error!")
}

// LogSuccess to be used to log success
func LogSuccess(cmd string, res string) {
	serverLogger.WithFields(logrus.Fields{
		"command": cmd,
		"result": res,
	}).Info("Query was successful!")
}
