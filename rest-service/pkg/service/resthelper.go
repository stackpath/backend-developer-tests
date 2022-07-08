package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func isValidNameQuery(c *gin.Context) bool {
	fmt.Println("Inside the Valid Query section")
	if _, ok := c.GetQuery("first_name"); ok {
		if _, ok := c.GetQuery("last_name"); ok {
			return true
		}
	}
	return false
}

func getName(c *gin.Context) (string, string) {
	if firstname, ok := c.GetQuery("first_name"); ok && len(firstname) > 0 {
		if lastname, ok := c.GetQuery("last_name"); ok && len(lastname) > 0 {
			return firstname, lastname
		}
	}
	return "", ""
}

func isValidPhoneQuery(c *gin.Context) bool {
	if _, ok := c.GetQuery("phone_number"); ok {
		return true
	}
	return false
}

func getPhoneNumber(c *gin.Context) string {
	if phonenumber, ok := c.GetQuery("phone_number"); ok && len(phonenumber) > 0 {
		return phonenumber
	}
	return ""
}
