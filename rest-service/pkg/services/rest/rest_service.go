package rest

import (
	servicedefinitions "github.com/stackpath/backend-developer-tests/rest-service/pkg/services/definitions"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/services/rest/people"
)

// Rest represents an implementation of a REST service
type Rest struct {
	People servicedefinitions.PeopleService
}

// New creates a new rest service instance
func New() (*Rest, error) {
	return &Rest{
		People: people.New(),
	}, nil
}
