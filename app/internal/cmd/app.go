package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/steelWinds/animals-rest-go/internal/app/animals"
	"github.com/steelWinds/animals-rest-go/internal/app/database"
	"github.com/steelWinds/animals-rest-go/internal/app/owners"
	"gorm.io/gorm"
)

func InitApp() (engine *gin.Engine, db *gorm.DB) {
	engine = gin.Default()

	var connErr error

	if db, connErr = database.ConnectDB(); connErr != nil {
		log.Fatalln(connErr.Error())
	}

	v1 := engine.Group("/api/v1")
	{
		animals.RegisterHandlers(v1, animals.NewService(db))
		owners.RegisterHandlers(v1, owners.NewService(db))
	}

	return
}