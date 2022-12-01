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

	items := make([]db.OwnerUnit, 2)

	orm.Model(new(db.OwnerUnit)).Preload("AnimalUnits").Find(&items)

	router.GET("/animals", handlers.GetH[db.AnimalUnit](orm))
	router.POST("/animals", handlers.PostH[db.AnimalUnit](orm))

	router.GET("/owners", handlers.GetH[db.OwnerUnit](orm, "AnimalUnits"))

	router.Run(":3000")
}
