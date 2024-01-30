package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func getJson(filepath string) {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Successfully opened PersonalData.json")

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteValue, &people)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if jsonErr := jsonFile.Close(); jsonErr != nil {
			fmt.Println("Can't close : ", jsonErr)
			return
		}
	}()
}

func getPeople(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, people)
}

func postPerson(c *gin.Context) {
	var newPerson person

	if err := c.BindJSON(&newPerson); err != nil {
		return
	}

	// Add the new album to the slice.
	people = append(people, newPerson)
	c.IndentedJSON(http.StatusCreated, newPerson)
}

func getPersonByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range people {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func getPersonByIndex(c *gin.Context) {
	index, _ := strconv.Atoi(c.Param("index"))

	for _, a := range people {
		if a.Index == index {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func getPersonByGUID(c *gin.Context) {
	guid := c.Param("guid")

	for _, a := range people {
		if a.GUID == guid {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func getPersonByIsActive(c *gin.Context) {
	isActive, _ := strconv.ParseBool(c.Param("isActive"))

	for _, a := range people {
		if a.IsActive == isActive {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func getPersonByBalance(c *gin.Context) {
	balance := c.Param("balance")

	for _, a := range people {
		if a.Balance == balance {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func getPersonByAge(c *gin.Context) {
	age, _ := strconv.Atoi(c.Param("isActive"))

	for _, a := range people {
		if a.Age == age {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func getPersonByEyeColor(c *gin.Context) {
	eyeColor := c.Param("eyeColor")

	for _, a := range people {
		if a.EyeColor == eyeColor {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

//TODO Person by Name

func getPersonByGender(c *gin.Context) {
	gender := c.Param("gender")

	for _, a := range people {
		if a.Gender == gender {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func getPersonByCompany(c *gin.Context) {
	company := c.Param("company")

	for _, a := range people {
		if a.Company == company {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func getPersonByEmail(c *gin.Context) {
	email := c.Param("isActive")

	for _, a := range people {
		if a.Email == email {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func getPersonByPhoneNumber(c *gin.Context) {
	phoneNumber := c.Param("phone")

	for _, a := range people {
		if a.Phone == phoneNumber {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

//TODO Person by address

func getPersonByAbout(c *gin.Context) {
	about := c.Param("about")

	for _, a := range people {
		if a.About == about {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func getPersonByRegistered(c *gin.Context) {
	registered := c.Param("registered")

	for _, a := range people {
		if a.Registered == registered {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func getPersonByLatitude(c *gin.Context) {
	latitude, _ := strconv.ParseFloat(c.Param("latitude"), 64)

	for _, a := range people {
		if a.Latitude == latitude {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func getPersonByLongitude(c *gin.Context) {
	longitude, _ := strconv.ParseFloat(c.Param("longitude"), 64)

	for _, a := range people {
		if a.Longitude == longitude {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

//TODO Person by Friends
