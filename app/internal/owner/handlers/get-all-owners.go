package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelWinds/animals-rest-go/internal/database"
	"github.com/steelWinds/animals-rest-go/internal/models"
	"gorm.io/gorm/clause"
)


func GetAllOwners(c *gin.Context) {
	var owners []models.OwnerUnit

	orm := database.GetDB().Model(new(models.OwnerUnit))

	if tx := orm.Preload(clause.Associations).Find(&owners); tx.Error != nil {
		c.JSON(http.StatusInternalServerError, tx.Error.Error())
	} 

	c.JSON(http.StatusOK, &owners)
}