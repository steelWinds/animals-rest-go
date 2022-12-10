package tests

import (
	"context"
	"log"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/steelWinds/animals-rest-go/internal/app/animals"
	"github.com/steelWinds/animals-rest-go/internal/app/owners"
	"github.com/steelWinds/animals-rest-go/internal/app/tests/resources"
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

	AnimalsService *animals.AnimalsSet
	OwnersService *owners.OwnersSet
}

type Message = map[string]string

func (suite *AppSuite) SetupSuite() {
	if err := godotenv.Load("resources/.env"); err != nil {
		suite.T().Fatal(err)
	}

	suite.App, suite.DB = cmd.InitApp()

	suite.AnimalsService = animals.NewService(suite.DB)
	suite.OwnersService = owners.NewService(suite.DB)
}

func TestSuite(t *testing.T) {
	suiteApp := new(AppSuite)

	compose, err := tc.NewDockerCompose("resources/docker-compose.test.yml")

	if err != nil {
		log.Fatal(err)
	}

	t.Cleanup(func() {
		db, err := suiteApp.DB.DB()

		if err != nil {
			log.Fatal(err)
		}

		db.Close()
	})

	ctx, cancel := context.WithCancel(context.Background())

	t.Cleanup(cancel)

	compose.
		WaitForService("postgres", wait.ForHealthCheck()).
		Up(ctx, tc.Wait(true));

	t.Cleanup(func () {
		resources.Clear(suiteApp.DB)
		compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal)
	})


	suite.Run(t, suiteApp)
}
