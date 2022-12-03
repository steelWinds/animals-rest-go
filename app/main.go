package main

import (
	"github.com/gin-gonic/gin"

	"github.com/steelWinds/animals-rest-go/db"
	"github.com/steelWinds/animals-rest-go/handlers"
)

func main() {
	router := gin.Default()

	orm, err := db.ConnectDB("user", "secret", "animals", "disable", 3001)

	if (err != nil) {
		panic(err)
	}

	v1 := router.Group("/api/v1")
	{
		v1.GET("/animals", handlers.GetH[db.AnimalUnit](orm, "OwnerUnits"))
		v1.POST("/animals", handlers.PostH[db.AnimalUnit](orm))

		v1.GET("/animals/:id", handlers.GetByIdH[db.AnimalUnit](orm, "OwnerUnits"))

		v1.GET("/owners", handlers.GetH[db.OwnerUnit](orm, "AnimalUnits"))
		v1.POST("/owners", handlers.PostH[db.OwnerUnit](orm))

		v1.POST("/owners/:id", handlers.PostAnimlsByIdH(orm))
		v1.GET("/owners/:id", handlers.GetByIdH[db.OwnerUnit](orm, "AnimalUnits"))
	}

	router.Run(":3000")
}
