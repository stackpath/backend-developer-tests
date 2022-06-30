package people

import (
	servicedefinitions "github.com/stackpath/backend-developer-tests/rest-service/pkg/services/definitions"
)

// People defines people service instance
type People struct{}

// New creates and returns a new user service instance
func New() servicedefinitions.PeopleService {
	return &People{}
}
