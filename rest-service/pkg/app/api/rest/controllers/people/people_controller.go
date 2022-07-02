package people

import (
	"github.com/gin-gonic/gin"
	servicedefinitions "github.com/stackpath/backend-developer-tests/rest-service/pkg/services/definitions"
)

// PeopleController is people REST API controller
type PeopleController struct {
	service servicedefinitions.PeopleService
}

// New creates and returns a user controller instance
func New(service servicedefinitions.PeopleService, router *gin.RouterGroup) *PeopleController {
	c := &PeopleController{service: service}

	router.GET("/people/:person_id", c.GetPerson)
	router.GET("/people", c.ListPeople)

	return c
}
