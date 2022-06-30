package people

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/log"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
)

// GetPerson returns a single user information
func (p *People) GetPerson(ctx *gin.Context, id uuid.UUID) (*models.Person, error) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	logger.Info("Get people service method called")

	person, err := models.FindPersonByID(id)
	if err != nil {
		return nil, err
	}
	return person, nil
}
