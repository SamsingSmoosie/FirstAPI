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
