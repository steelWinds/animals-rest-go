package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelWinds/animals-rest-go/internal/database"
	dbModels "github.com/steelWinds/animals-rest-go/internal/database/models"
	"github.com/steelWinds/animals-rest-go/internal/models"
)


func GetOwnerById(c *gin.Context) {
	var itemParams dbModels.IDParams

	if err := c.ShouldBindUri(&itemParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })

		return
	}

	orm := database.GetDB().Model(new(models.OwnerUnit))

	var item models.OwnerUnit

	if tx := orm.Where("id = ?", itemParams.ID).Preload("AnimalUnits").First(&item); tx.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{ "message": tx.Error.Error() })

		return
	}

	c.IndentedJSON(http.StatusOK, &item)
}