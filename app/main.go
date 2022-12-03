package main

import (
	"github.com/gin-gonic/gin"

	animalH "github.com/steelWinds/animals-rest-go/internal/animal/handlers"
	"github.com/steelWinds/animals-rest-go/internal/database"
	otherH "github.com/steelWinds/animals-rest-go/internal/handlers"
	"github.com/steelWinds/animals-rest-go/internal/models"
	ownerH "github.com/steelWinds/animals-rest-go/internal/owner/handlers"
)

func main() {
	router := gin.Default()

	database.ConnectDB("animals", "disable")

	v1 := router.Group("/api/v1")
	{
		v1.GET("/animals", animalH.GetAllAnimals)
		v1.GET("/animal/:id", animalH.GetAnimalById)
		v1.POST("/animal", otherH.PostItem[models.AnimalUnit])

		v1.GET("/owners", ownerH.GetAllOwners)
		v1.GET("/owner/:id", ownerH.GetOwnerById)
		v1.POST("/owner", otherH.PostItem[models.OwnerUnit])
		v1.POST("/owner/:id", ownerH.PostAnimlsById)
	}

	router.Run(":3000")
}
