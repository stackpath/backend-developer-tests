package definitions

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
)

//go:generate mockery --name=PeopleService

// PeopleService defines a methods for people service
type PeopleService interface {
	GetPerson(ctx *gin.Context, id uuid.UUID) (*models.Person, error)
	ListPeople(ctx *gin.Context, filters *models.PeopleFilter) ([]*models.Person, error)
}
