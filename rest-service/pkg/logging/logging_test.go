package logging

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type LoggingSuite struct {
	suite.Suite
}

func TestLoggingSuite(t *testing.T) {
	suite.Run(t, new(LoggingSuite))
}

func (suite LoggingSuite) TestInitLogger() {
	Logger = log.New()
	InitLogger()
	suite.Equal("debug", Logger.GetLevel().String())

	suite.Equal(&log.JSONFormatter{}, Logger.Formatter)
}
