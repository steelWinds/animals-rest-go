package tests

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/steelWinds/animals-rest-go/internal/app/animals"
	"github.com/steelWinds/animals-rest-go/internal/app/models"
	"github.com/steelWinds/animals-rest-go/internal/cmd"
	"github.com/stretchr/testify/suite"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"
)

type AppSuite struct {
	suite.Suite
	App *gin.Engine
	DB *gorm.DB

	Service *animals.AnimalsSet
	CreatedItems []models.AnimalUnit
}

type Message = map[string]string

func (suite *AppSuite) SetupSuite() {
	var err error
	var db *sql.DB

	if err = godotenv.Load("resources/.test.env"); err != nil {
		suite.T().Fatal(err)
	}

	suite.App, suite.DB = cmd.InitApp()
	
	if db, err = suite.DB.DB(); err != nil {
		suite.T().Fatal(err)
	}

	if err = SetTestMigrations(db); err != nil {
		suite.T().Fatal(err)
	}

	suite.Service = animals.NewService(suite.DB)
}

func TestSuite(t *testing.T) {
	suiteApp := new(AppSuite)

	compose, err := tc.NewDockerCompose("resources/docker-compose.test.yml")

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	t.Cleanup(cancel)

	compose.
		WaitForService("postgres", wait.ForExposedPort()).
		Up(ctx, tc.Wait(true));


	t.Cleanup(func () {
		Clear(suiteApp.DB)
		compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal)
	})


	suite.Run(t, suiteApp)
}