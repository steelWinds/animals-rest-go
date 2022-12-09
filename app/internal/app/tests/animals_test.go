package tests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/lib/pq"
	"github.com/steelWinds/animals-rest-go/internal/app/models"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"gorm.io/gorm/clause"
)

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