package main

import (
	"github.com/gin-gonic/gin"
)

var people []person

func main() {

	getJson("./PersonalData.json")

	router := gin.Default()
	router.GET("/people", getPeople)
	router.GET("/albums/id/:id", getPersonByID)
	router.GET("/people/index/:index", getPersonByIndex)
	router.POST("/people", postPerson)

	_ = router.Run("localhost:8080")
}
