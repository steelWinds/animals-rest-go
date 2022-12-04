package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/steelWinds/animals-rest-go/internal/app/animals"
	"github.com/steelWinds/animals-rest-go/internal/app/database"
	"github.com/steelWinds/animals-rest-go/internal/app/owners"
)

func main() {
	router := gin.Default()

	db, err := database.ConnectDB("animals", "disable")

	if err != nil {
		log.Fatalln(err.Error())
	}

	v1 := router.Group("/api/v1")
	{
		animals.RegisterHandlers(v1, animals.NewService(db))
		owners.RegisterHandlers(v1, owners.NewService(db))
	}

	router.Run(":3000")
}
