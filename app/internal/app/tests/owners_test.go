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

func (suite *AppSuite) TestPostOwnerItem() {
	var item models.AnimalUnit

	gofakeit.Struct(&item)

	apitest.New().
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
}

func (suite *AppSuite) TestGetOwnerItemById() {
	var item models.OwnerUnit
	var createdItem models.OwnerUnit
	var err error

	gofakeit.Struct(&item)

	if createdItem, err = suite.OwnersService.CreateItem(item); err != nil {
		suite.T().Error(err)
	}

	reqStr := fmt.Sprintf("/api/v1/owners/%d", createdItem.ID)

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
		End(),
	).
	End()
}

func (suite *AppSuite) TestGetOwnersAllItems() {
	var list []models.OwnerUnit

	suite.DB.Preload(clause.Associations).Find(&list)

	apitest.New().
	HandlerFunc(suite.App.ServeHTTP).
	Get("/api/v1/owners").
	Expect(suite.T()).
	Status(http.StatusOK).
	Assert(jsonpath.GreaterThan("$", 0)).
	End()
}

func (suite *AppSuite) TestOwnerDelete() {
	var item models.OwnerUnit
	var createdItem models.OwnerUnit
	var err error
	var messageJSON []byte

	gofakeit.Struct(&item)

	if createdItem, err = suite.OwnersService.CreateItem(item); err != nil {
		suite.T().Error(err)
	}

	reqStr := fmt.Sprintf("/api/v1/owners/%d", createdItem.ID)

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
