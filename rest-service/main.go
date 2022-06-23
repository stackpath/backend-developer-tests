package main

import (
	"backend-developer-tests/rest-service/pkg/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	// TODO: Add RESTful web service here
	router := gin.Default()

	router.GET("/people", func(c *gin.Context) {
		if firstName, has := c.GetQuery("first_name"); has {
			lastName, _ := c.GetQuery("last_name")
			fmt.Println(firstName, lastName)
			persons := models.FindPeopleByName(firstName, lastName)
			c.JSON(http.StatusOK, persons)
			return
		}

		if phoneNumber, has := c.GetQuery("phone_number"); has {
			fmt.Println(phoneNumber)
			persons := models.FindPeopleByPhoneNumber(phoneNumber)
			c.JSON(http.StatusOK, persons)
			return
		}
		c.JSON(http.StatusOK, models.AllPeople())
	})

	router.GET("/people/:id", func(c *gin.Context) {
		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
		}
		if person, err := models.FindPersonByID(id); err != nil {
			c.JSON(http.StatusNotFound, nil)
		} else {
			c.JSON(http.StatusOK, person)
		}
	})
	router.Run()
}
