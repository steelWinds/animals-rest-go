package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/steelWinds/animals-rest-go/handlers"
)

func main() {
	fmt.Println("I'm first go-docker app!")

	router := gin.Default()

	router.GET("/animals", handlers.GetAnimals)
	router.POST("/animals", handlers.PostAnimal)

	router.Run(":3000")
}