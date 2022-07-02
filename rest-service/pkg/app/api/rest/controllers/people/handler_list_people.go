package people

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/log"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
)

// ListUsers lists users
func (p *PeopleController) ListPeople(ctx *gin.Context) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	logger.Info("List people handler called")

	firstName := ctx.Query("first_name")
	lastName := ctx.Query("last_name")
	phoneNumber := ctx.Query("phone_number")

	// validate inputs, not doing now (TBD)
	var filters *models.PeopleFilter
	if firstName != "" || lastName != "" || phoneNumber != "" {
		filters = &models.PeopleFilter{}
		filters.FirstName = firstName
		filters.LastName = lastName
		filters.PhoneNumber = phoneNumber
	}

	people, err := p.service.ListPeople(ctx, filters)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.RestErr{
			Code:   http.StatusNotFound,
			Errors: []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, people)
}
