package models

import (
	"encoding/json"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// Person defines a simple representation of a person
type Person struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
}

// people is the data source for the People RESTful service.
var people = []*Person{
	{
		ID:          uuid.Must(uuid.FromString("81eb745b-3aae-400b-959f-748fcafafd81")),
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "+1 (800) 555-1212",
	},
	{
		ID:          uuid.Must(uuid.FromString("5b81b629-9026-450d-8e46-da4f8c7bd513")),
		FirstName:   "Jane",
		LastName:    "Doe",
		PhoneNumber: "+1 (800) 555-1313",
	},
	{
		ID:          uuid.Must(uuid.FromString("df12ce76-767b-4bf0-bccb-816745df9e70")),
		FirstName:   "Brian",
		LastName:    "Smith",
		PhoneNumber: "+44 7700 900077",
	},
	// This is another person with the name John Doe
	{
		ID:          uuid.Must(uuid.FromString("135af595-aa86-4bb5-a8f7-df17e6148e63")),
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "+1 (800) 555-1414",
	},
	// This is another person with the phone number +44 7700 900077
	{
		ID:          uuid.Must(uuid.FromString("000ebe58-b659-422b-ab48-a0d0d40bd8f9")),
		FirstName:   "Jenny",
		LastName:    "Smith",
		PhoneNumber: "+44 7700 900077",
	},
}

// AllPeople returns all people in `people`.
func AllPeople() []*Person {
	return people
}

// FindPersonByID searches for people in `people` the by their ID.
func FindPersonByID(id uuid.UUID) (*Person, error) {
	for _, person := range people {
		if person.ID == id {
			return person, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("user ID %s not found", id.String()))
}

// FindPeopleByName performs a case-sensitive search for people in `people` by first or last name.
func FindPeopleByName(firstName, lastName string) []*Person {
	result := make([]*Person, 0)

	for _, person := range people {
		if firstName != "" && person.FirstName != firstName {
			continue
		}
		if lastName != "" && person.LastName != lastName {
			continue
		}

		result = append(result, person)
	}

	return result
}

// FindPeopleByPhoneNumber searches for people in `people` by phone number.
func FindPeopleByPhoneNumber(phoneNumber string) []*Person {
	result := make([]*Person, 0)

	for _, person := range people {
		if person.PhoneNumber == phoneNumber {
			result = append(result, person)
		}
	}

	return result
}

// ToJSON represents a person as a JSON string.
func (person *Person) ToJSON() (string, error) {
	marshaled, err := json.Marshal(person)
	if err != nil {
		return "", err
	}

	return string(marshaled[:]), nil
}
