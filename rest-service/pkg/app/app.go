package app

import "backend-developer-tests/rest-service/pkg/models"

func GetPeople() []*models.Person {
	return models.AllPeople()
}
