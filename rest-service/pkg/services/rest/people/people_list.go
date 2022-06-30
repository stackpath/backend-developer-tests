package people

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/log"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
)

// ListPeople returns list of people information
func (p *People) ListPeople(ctx *gin.Context, filters *models.PeopleFilter) ([]*models.Person, error) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	logger.Info("List people service method called")

	if filters == nil {
		return models.AllPeople(), nil
	}

	people := make([]*models.Person, 0)
	if filters.FirstName != "" || filters.LastName != "" {
		people = models.FindPeopleByName(filters.FirstName, filters.LastName)
		if len(people) == 0 {
			return people, fmt.Errorf(
				"People information with first name %s and last name %s not found",
				filters.FirstName,
				filters.LastName,
			)
		}
	}

	if filters.PhoneNumber != "" {
		people = models.FindPeopleByPhoneNumber(filters.PhoneNumber)
		if len(people) == 0 {
			return people, fmt.Errorf(
				"People information with phone number %s not found",
				filters.PhoneNumber,
			)
		}
	}

	return people, nil
}
