package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
)

func GetAllPeople(c *gin.Context) {

	if isValidNameQuery(c) && isValidPhoneQuery(c) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "query with name and phone is not supported"})
		return
	}

	if isValidNameQuery(c) {
		firstname, lastname := getName(c)
		var message string

		if len(firstname) == 0 && len(lastname) == 0 {
			message = "Firstname and Lastname missing in the query"
		} else if len(firstname) == 0 {
			message = "Firstname missing in the query"
		} else if len(lastname) == 0 {
			message = "Lastname missing in the query"
		}

		if len(message) != 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": message})
			return
		}
		c.IndentedJSON(http.StatusOK, models.FindPeopleByName(firstname, lastname))
		return

	} else if isValidPhoneQuery(c) {
		phonenumber := getPhoneNumber(c)
		if len(phonenumber) == 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "phone number missing in the query"})
			return
		}
		fmt.Println("Phone number is ", phonenumber)

		c.IndentedJSON(http.StatusOK, models.FindPeopleByPhoneNumber(phonenumber))
		return
	}

	c.IndentedJSON(http.StatusOK, models.AllPeople())
}

func FindPersonByID(c *gin.Context) {
	id := c.Param("id")
	if person, err := models.FindPersonByID(uuid.Must(uuid.Parse(id))); err == nil {
		c.IndentedJSON(http.StatusOK, person)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
	}
}
