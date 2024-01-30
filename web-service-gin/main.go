package main

import (
	"github.com/gin-gonic/gin"
)

var people []person

func main() {

	getJson("./PersonalData.json")

	router := gin.Default()
	router.GET("/people", getPeople)
	router.GET("/people/id/:id", getPersonByID)
	router.GET("/people/index/:index", getPersonByIndex)
	router.GET("/people/guid/:guid", getPersonByGUID)
	router.GET("/people/isActive/:isActive", getPersonByIsActive)
	router.GET("/people/balance/:balance", getPersonByBalance)
	router.GET("/people/age/:age", getPersonByAge)
	router.GET("/people/eyeColor/:eyeColor", getPersonByEyeColor)
	//TODO router.GET("/people/name/:name", getPersonByName)
	router.GET("/people/gender/:gender", getPersonByGender)
	router.GET("/people/company/:company", getPersonByCompany)
	router.GET("/people/email/:email", getPersonByEmail)
	router.GET("/people/phoneNumber/:phone", getPersonByPhoneNumber)
	//TODO router.GET("/people/address/:address", getPersonByAddress)
	router.GET("/people/about/:about", getPersonByAbout)
	router.GET("/people/registered/:registered", getPersonByRegistered)
	router.GET("/people/latitude/:latitude", getPersonByLatitude)
	router.GET("/people/longitude/:longitude", getPersonByLongitude)
	// TODO router.GET("/people/friends/:friends", getPersonByFriends)
	router.POST("/people", postPerson)

	_ = router.Run("localhost:8080")
}
