package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/steelWinds/animals-rest-go/internal/app/animals"
	"github.com/steelWinds/animals-rest-go/internal/app/database"
	"github.com/steelWinds/animals-rest-go/internal/app/models"
	"github.com/steelWinds/animals-rest-go/internal/cmd"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AppSuite struct {
	suite.Suite
	App *gin.Engine
	DB *gorm.DB

	Service *animals.AnimalsSet
	CreatedItems []models.AnimalUnit
}

type Message = map[string]string

func clearTables(suite *AppSuite) {
	var tables []string
	var err error

	if tables, err = suite.DB.Migrator().GetTables(); err != nil {
		suite.T().Fatal(err)
	}

	for _, table := range tables {
		suite.DB.Migrator().DropTable(table)
	}
}

func (suite *AppSuite) SetupSuite() {
	var err error
	var db *sql.DB

	if err = godotenv.Load(); err != nil {
		suite.T().Fatal(err)
	}

	suite.App, suite.DB = cmd.InitApp()
	
	if db, err = suite.DB.DB(); err != nil {
		suite.T().Fatal(err)
	}

	if err = database.SetTestMigrations(db); err != nil {
		suite.T().Fatal(err)
	}

	suite.Service = animals.NewService(suite.DB)
}

func (suite *AppSuite) TestPostItem() {
	var item models.AnimalUnit
	
	gofakeit.Struct(&item)

	result := apitest.New().
	HandlerFunc(suite.App.ServeHTTP).
	Post("/api/v1/animals").
	JSON(item).
	Expect(suite.T()).
	Status(http.StatusCreated).
	Assert(
		jsonpath.
		Chain().
		Equal("name", item.Name).
		Equal("age", float64(item.Age)).
		Equal("type", item.Type).
		End(),
	).
	End()

	result.JSON(&item)

	suite.CreatedItems = append(suite.CreatedItems, item)
}

func (suite *AppSuite) TestGetItemById() {
	var item models.AnimalUnit
	var createdItem models.AnimalUnit 
	var err error
	
	gofakeit.Struct(&item)

	if createdItem, err = suite.Service.CreateItem(item); err != nil {
		suite.T().Error(err)
	}

	reqStr := fmt.Sprintf("/api/v1/animals/%d", createdItem.ID)

	apitest.New().
	HandlerFunc(suite.App.ServeHTTP).
	Get(reqStr).
	JSON(item).
	Expect(suite.T()).
	Status(http.StatusOK).
	Assert(
		jsonpath.
		Chain().
		Equal("name", createdItem.Name).
		Equal("age", float64(createdItem.Age)).
		Equal("type", createdItem.Type).
		End(),
	).
	End()

	suite.CreatedItems = append(suite.CreatedItems, createdItem)
}

func (suite *AppSuite) TestGetAllItems() {
	var list []models.AnimalUnit

	suite.DB.Preload(clause.Associations).Find(&list)

	apitest.New().
	HandlerFunc(suite.App.ServeHTTP).
	Get("/api/v1/animals").
	Expect(suite.T()).
	Status(http.StatusOK).
	Assert(jsonpath.GreaterThan("$", 0)).
	End()
}

func (suite *AppSuite) TestDelete() {
	var item models.AnimalUnit
	var createdItem models.AnimalUnit
	var err error
	var messageJSON []byte
	
	gofakeit.Struct(&item)

	if createdItem, err = suite.Service.CreateItem(item); err != nil {
		suite.T().Error(err)
	}

	reqStr := fmt.Sprintf("/api/v1/animals/%d", createdItem.ID)

	if messageJSON, err = json.Marshal(Message{"message": "succesfull"}); err != nil {
		suite.T().Error(err)
	}

	apitest.New().
	HandlerFunc(suite.App.ServeHTTP).
	Delete(reqStr).
	Expect(suite.T()).
	Status(http.StatusOK).
	Body(string(messageJSON)).
	End()
}

func TestSuite(t *testing.T) {
	suiteApp := new(AppSuite)

	defer clearTables(suiteApp)

	suite.Run(t, suiteApp)
}