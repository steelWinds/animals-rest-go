package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelWinds/animals-rest-go/internal/database"
	"github.com/steelWinds/animals-rest-go/internal/models"
)


func GetAllAnimals(c *gin.Context) {
	var animals []models.AnimalUnit

	orm := database.GetDB().Model(new(models.AnimalUnit))

	if tx := orm.Preload("OwnerUnits").Find(&animals); tx.Error != nil {
		c.JSON(http.StatusInternalServerError, tx.Error.Error())
	} 

	c.JSON(http.StatusOK, &animals)
}