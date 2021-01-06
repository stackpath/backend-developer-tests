package logging

import (
	log "github.com/sirupsen/logrus"
)

// Logger is an instance of logrus logger
var Logger = log.New()

// InitLogger sets up logging level, and log formatting
func InitLogger() {
	Logger.SetLevel(log.DebugLevel)
	Logger.SetFormatter(&log.JSONFormatter{})
}
