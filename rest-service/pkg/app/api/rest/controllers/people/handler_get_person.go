package people

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/log"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
)

// GetPerson fetches a single person information
func (p *PeopleController) GetPerson(ctx *gin.Context) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	logger.Info("Get person handler called")

	id := ctx.Param("person_id")
	personID, err := uuid.FromString(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.RestErr{
			Code:   http.StatusBadRequest,
			Errors: []string{err.Error()},
		})
		return
	}

	person, err := p.service.GetPerson(ctx, personID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.RestErr{
			Code:   http.StatusNotFound,
			Errors: []string{err.Error()},
		})
		return
	}
	ctx.JSON(http.StatusOK, person)
}
