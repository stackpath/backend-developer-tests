package router

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RouterSuite struct {
	suite.Suite
}

func TestRouterSuite(t *testing.T) {
	suite.Run(t, new(RouterSuite))
}

func (suite RouterSuite) TestNewRouter() {
	r := NewRouter()
	routes := r.Routes()
	suite.Equal(routes[0].Path, "/people")
	suite.Equal(routes[0].Method, "GET")
	suite.Equal(routes[1].Path, "/people/:id")
	suite.Equal(routes[1].Method, "GET")
}
