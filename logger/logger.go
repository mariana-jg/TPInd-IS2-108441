package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLogger() {

	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Logger.Fatal("Failed to open log file:", err)
	}

	Logger.SetOutput(logFile)
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetLevel(logrus.InfoLevel)

}
