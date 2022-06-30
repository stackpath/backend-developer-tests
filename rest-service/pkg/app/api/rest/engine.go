package rest

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	peoplecontroller "github.com/stackpath/backend-developer-tests/rest-service/pkg/app/api/rest/controllers/people"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/app/api/rest/middlewares"
	restservice "github.com/stackpath/backend-developer-tests/rest-service/pkg/services/rest"
)

// Rest defines a REST application
type Rest struct {
	Engine *gin.Engine
}

// New creates a new REST application
func New(restService *restservice.Rest) (*Rest, error) {
	engine := gin.New()
	engine.Use(
		gin.Recovery(), // recover from panic
	)

	// base router group
	routes := engine.Group("")

	// add additional middlewares to base routes
	routes.Use(
		middlewares.RequestID(), // this should be before logging middleware
		middlewares.Logging(),
	)

	// root path to return OK response code
	routes.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	peopleRouterGroup := routes.Group("")
	peoplecontroller.New(restService.People, peopleRouterGroup)

	return &Rest{
		Engine: engine,
	}, nil
}

// Run starts a new REST listner
func (r *Rest) Run() error {
	return r.Engine.Run(
		net.JoinHostPort("0.0.0.0", "8080"),
	)
}
