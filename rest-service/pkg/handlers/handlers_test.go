// build +test

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
	"github.com/stretchr/testify/suite"
)

type HandlerSuite struct {
	suite.Suite
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (suite HandlerSuite) TestGetPeopleHandler() {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/people", nil)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(request, recorder)
	ctx.SetPath("people")

	if suite.NoError(GetPeopleHandler(ctx)) {
		var people []models.Person
		err := json.NewDecoder(recorder.Body).Decode(&people)
		suite.NoError(err)
		suite.Equal(recorder.Code, http.StatusOK)
		suite.Equal(people[0].FirstName, "John")
		suite.Equal(people[0].LastName, "Doe")
		suite.Equal(len(people), 5)
	}
}

func (suite HandlerSuite) TestGetPersonByIDHandler() {
	cases := []struct {
		id        string
		shouldErr bool
		errStatus int
	}{
		{
			id: "81eb745b-3aae-400b-959f-748fcafafd81",
		},
		{
			id:        "notauuid",
			shouldErr: true,
			errStatus: http.StatusBadRequest,
		},
		{
			id:        "0ce61c17-5eb4-4587-8c36-dcf4062ada4c",
			shouldErr: true,
			errStatus: http.StatusNotFound,
		},
	}

	for i, c := range cases {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/people/%s", c.id), nil)
		recorder := httptest.NewRecorder()
		ctx := e.NewContext(request, recorder)
		ctx.SetPath("/people/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues(c.id)

		_ = GetPersonByIDHandler(ctx)
		if c.shouldErr {
			suite.Equal(c.errStatus, recorder.Code, "Test case: %d", i)
		} else {
			var person models.Person
			err := json.NewDecoder(recorder.Body).Decode(&person)
			suite.NoError(err, "Test case: %d", i)
			suite.Equal(http.StatusOK, recorder.Code, "Test case: %d", i)
		}
	}
}

func (suite HandlerSuite) TestGetPeopleWithName() {
	cases := []struct {
		firstName           string
		lastName            string
		totalPeopleExpected int
		shouldErr           bool
		errStatus           int
	}{
		{
			firstName:           "John",
			lastName:            "Doe",
			totalPeopleExpected: 2,
		},
		{
			firstName: "onlyfirstnameprovided",
			shouldErr: true,
			errStatus: http.StatusBadRequest,
		},
		{
			lastName:  "onlylastnameprovided",
			shouldErr: true,
			errStatus: http.StatusBadRequest,
		},
		{
			firstName:           "notfound",
			lastName:            "notfound",
			errStatus:           http.StatusOK,
			totalPeopleExpected: 0,
		},
	}

	for i, c := range cases {
		e := echo.New()
		q := make(url.Values)
		q.Set("first_name", c.firstName)
		q.Set("last_name", c.lastName)
		request := httptest.NewRequest(http.MethodGet, "/people?"+q.Encode(), nil)
		recorder := httptest.NewRecorder()
		ctx := e.NewContext(request, recorder)
		ctx.SetPath("people")

		fmt.Println(ctx.Request().URL)

		_ = GetPeopleHandler(ctx)
		if c.shouldErr {
			suite.Equal(c.errStatus, recorder.Code, "Test case: %d", i)
		} else {
			var people []models.Person
			err := json.NewDecoder(recorder.Body).Decode(&people)
			suite.NoError(err, "Test case: %d", i)

			suite.Equal(recorder.Code, http.StatusOK, "Test case: %d", i)
			suite.Equal(c.totalPeopleExpected, len(people), "Test case: %d", i)
		}
	}
}

func (suite HandlerSuite) TestGetPeopleByPhone() {
	cases := []struct {
		phoneNumber         string
		totalPeopleExpected int
		shouldErr           bool
		errStatus           int
	}{
		{
			phoneNumber:         "+1 (800) 555-1212",
			totalPeopleExpected: 1,
		},
		{
			phoneNumber: "notavalidphonenumber",
			shouldErr:   true,
			errStatus:   http.StatusBadRequest,
		},
		{
			phoneNumber:         "+1 (111) 111-1111",
			errStatus:           http.StatusOK,
			totalPeopleExpected: 0,
		},
	}

	for i, c := range cases {
		e := echo.New()
		q := make(url.Values)
		q.Set("phone_number", c.phoneNumber)
		request := httptest.NewRequest(http.MethodGet, "/people?"+q.Encode(), nil)
		recorder := httptest.NewRecorder()
		ctx := e.NewContext(request, recorder)
		ctx.SetPath("people")

		fmt.Println(ctx.Request().URL)

		_ = GetPeopleHandler(ctx)
		if c.shouldErr {
			suite.Equal(c.errStatus, recorder.Code, "Test case: %d", i)
		} else {
			var people []models.Person
			err := json.NewDecoder(recorder.Body).Decode(&people)
			suite.NoError(err, "Test case: %d", i)

			suite.Equal(recorder.Code, http.StatusOK, "Test case: %d", i)
			suite.Equal(c.totalPeopleExpected, len(people), "Test case: %d", i)
		}
	}
}
