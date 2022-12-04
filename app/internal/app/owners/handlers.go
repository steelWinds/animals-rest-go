package owners

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelWinds/animals-rest-go/internal/app/models"
)

func RegisterHandlers(router *gin.RouterGroup, service *OwnersSet) {
	res := resource{service} 

	router.GET("/owner/:id", res.get)
	router.POST("/owner/:id", res.postAnimals)
	router.GET("/owners", res.getAll)
	
	router.POST("/owners", res.post)
}

type resource struct {
	*OwnersSet
}

func (res *resource) get(c *gin.Context) {
	var params models.IDParam

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })

		return
	}

	owner, err := res.GetItem(params.ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{ "message": err.Error() })

		return
	}

	c.JSON(http.StatusOK, &owner)
}

func (res *resource) getAll(c *gin.Context) {
	animals, err := res.GetAllItems()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} 

	c.JSON(http.StatusOK, &animals)
}

func (res *resource) post(c *gin.Context) {
	var owner models.OwnerUnit

	if err := c.BindJSON(&owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	createdAnimal, err := res.CreateItem(owner)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusCreated, &createdAnimal)
}

func (res *resource) postAnimals(c *gin.Context) {
	var params models.IDParam

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })

		return
	}

	var animals []models.AnimalUnit

	if err := c.BindJSON(&animals); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })

		return
	}

	owner, err := res.PostAnimals(params.ID, animals)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error() })

		return
	}

	c.JSON(http.StatusCreated, &owner)
}