package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelWinds/animals-rest-go/db"
)

func GetAnimals(c *gin.Context) {
	postgres, err := db.ConnectDB("user", "secret", "animals", "disable", 3001)

	if (err != nil) {
		panic(err)
	}

	animals := make([]db.AnimalUnit, 0)

	postgres.Select(&animals, "SELECT * FROM animals_units")

	c.IndentedJSON(http.StatusOK, &animals)
}