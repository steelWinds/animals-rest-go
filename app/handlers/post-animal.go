package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelWinds/animals-rest-go/db"
)

func PostAnimal(c *gin.Context) {
	postgres, err := db.ConnectDB("user", "secret", "animals", "disable", 3001)

	if (err != nil) {
		panic(err)
	}

	animal := new(db.AnimalUnit)

	if err := c.BindJSON(&animal); err != nil {
		return
	}

	queryString := `
		INSERT INTO animals_units(name, age, type)
		VALUES ($1, $2, $3)
		RETURNING animal_id
	`

	createdErr := postgres.QueryRowx(queryString, animal.Name, animal.Age, animal.Type).Scan(&animal.AnimalID)

	if createdErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{ "message": createdErr.Error() })

		return
	}

	c.IndentedJSON(http.StatusOK, &animal)
}
