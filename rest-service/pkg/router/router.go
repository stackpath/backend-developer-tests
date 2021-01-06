package router

import (
	"github.com/labstack/echo/v4"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/handlers"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	// Setup Route Handlers
	e.GET("/people", handlers.GetPeopleHandler)
	e.GET("/people/:id", handlers.GetPersonByIDHandler)
	return e
}
