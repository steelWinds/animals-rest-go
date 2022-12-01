package main

import (
	"github.com/gin-gonic/gin"

	"github.com/steelWinds/animals-rest-go/handlers"
)

func main() {
	router := gin.Default()

	router.GET("/animals", handlers.GetAnimals)
	router.POST("/animals", handlers.PostAnimal)

	router.Run(":3000")
}
