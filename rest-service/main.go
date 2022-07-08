package main

import (
	"fmt"

	"github.com/stackpath/backend-developer-tests/rest-service/pkg/service"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println()

	router := gin.Default()
	router.GET("/people", service.GetAllPeople)
	router.GET("/people/:id", service.FindPersonByID)
	router.Run("localhost:8080")
}
