package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

var person = []Person{
	{1, "sandesh", "Belapur CBD"},
	{2, "Vishal", "At. Post Ulwe"},
}

func getPerson(c *gin.Context) {
	c.JSON(http.StatusOK, person)

}

func main() {
	router := gin.Default()

	router.GET("/person", getPerson)
	router.Run(":8080")
}
