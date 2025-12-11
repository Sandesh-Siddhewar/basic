package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

var person = []Person{
	{"1", "sandesh", "Belapur CBD"},
	{"2", "Vishal", "At. Post Ulwe"},
}

func getPersons(c *gin.Context) {
	c.JSON(http.StatusOK, person)

}

func getPerson(c *gin.Context) {
	id := c.Param("id")

	for _, person := range person {
		if person.Id == id {
			c.JSON(http.StatusOK, person)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"Message": "Person not found"})
}

func createPerson(c *gin.Context) {
	var newPerson Person

	if err := c.BindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Person := append(person, newPerson)
	c.JSON(http.StatusCreated, Person)

}
func main() {
	router := gin.Default()

	router.GET("/persons", getPersons)
	router.GET("/person/:id", getPerson)
	router.Run(":8081")
}
