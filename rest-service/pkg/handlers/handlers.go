package handlers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/logging"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
)

// GetPeople will retrieve a list containing all people in the system
func GetPeopleHandler(c echo.Context) error {
	firstName := c.QueryParam("first_name")
	lastName := c.QueryParam("last_name")
	phoneNumber := c.QueryParam("phone_number")

	// Handle requests with name, or phone query params
	_ = handleQueryPeopleByName(c, firstName, lastName)
	_ = handleQueryPeopleByPhoneNumber(c, phoneNumber)

	people := models.AllPeople()
	return c.JSON(http.StatusOK, people)
}

// GetPersonByID will retrieve a person with the provided ID
func GetPersonByIDHandler(c echo.Context) error {
	id := c.Param("id")
	uuidFromString, err := uuid.FromString(id)
	if err != nil {
		logging.Logger.Errorf("Person ID must be in uuid format. Received: %s", id)
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Person ID must be in uuid format. Received: %s", id))
	}

	person, err := models.FindPersonByID(uuidFromString)

	if err != nil {
		// Only error case from this method is not found
		logging.Logger.Errorf("Error: unable to find person with id: %s", id)
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, person)
}

func handleQueryPeopleByName(c echo.Context, firstName, lastName string) error {
	if len(firstName) > 0 || len(lastName) > 0 {
		logging.Logger.Info("Processing request: Query people by first & last name")
		if firstName == "" {
			logging.Logger.Errorf("Error: first_name query parameter is empty when querying people by name.")
			return c.JSON(http.StatusBadRequest, "Please provide a first name when querying people by name.")
		}

		if lastName == "" {
			logging.Logger.Errorf("Error: last_name query parameter is empty when querying people by name.")
			return c.JSON(http.StatusBadRequest, "Please provide a last name when querying people by name.")
		}

		people := models.FindPeopleByName(firstName, lastName)
		return c.JSON(http.StatusOK, people)
	}
	return nil
}

func handleQueryPeopleByPhoneNumber(c echo.Context, phoneNumber string) error {
	if len(phoneNumber) > 0 {
		pattern := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
		if !pattern.MatchString(phoneNumber) {
			logging.Logger.Errorf("Error: Invalid phone number. Received: %s", phoneNumber)
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("phone_number query must be a valid phone number format. Received: %s", phoneNumber))
		}
		people := models.FindPeopleByPhoneNumber(phoneNumber)
		return c.JSON(http.StatusOK, people)
	}
	return nil
}
