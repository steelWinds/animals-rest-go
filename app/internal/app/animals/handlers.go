package animals

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelWinds/animals-rest-go/internal/app/models"
)

func RegisterHandlers(router *gin.RouterGroup, service *AnimalsSet) {
	res := resource{service} 

	router.GET("/animals/:id", res.get)
	router.GET("/animals", res.getAll)
	
	router.POST("/animals", res.post)
}

type resource struct {
	*AnimalsSet
}

func (res *resource) get(c *gin.Context) {
	var params models.IDParam

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })

		return
	}

	animal, err := res.GetItem(params.ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{ "message": err.Error() })

		return
	}

	c.JSON(http.StatusOK, &animal)
}

func (res *resource) getAll(c *gin.Context) {
	animals, err := res.GetAllItems()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} 

	c.JSON(http.StatusOK, &animals)
}

func (res *resource) post(c *gin.Context) {
	var animal models.AnimalUnit

	if err := c.BindJSON(&animal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	createdAnimal, err := res.CreateItem(animal)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusCreated, &createdAnimal)
}