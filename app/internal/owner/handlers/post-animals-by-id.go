package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelWinds/animals-rest-go/internal/database"
	dbModels "github.com/steelWinds/animals-rest-go/internal/database/models"
	"github.com/steelWinds/animals-rest-go/internal/models"
)

func PostAnimlsById(c *gin.Context) {
	var ownerParams dbModels.IDParams

	if err := c.ShouldBindUri(&ownerParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })

		return
	}

	var animals []models.AnimalUnit

	if err := c.BindJSON(&animals); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })

		return
	}

	var ownMappers []models.OwnMap
	
	for _, animal := range animals {
		ownMappers = append(
			ownMappers, 
			models.OwnMap{
				AnimalUnitID: animal.ID,
				OwnerUnitID: ownerParams.ID,
			},
		)
	}

	orm := database.GetDB()

	if tx := orm.Create(&ownMappers); tx.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"message": tx.Error.Error()})

		return
	}

	var ownersAnimals []models.AnimalUnit

	orm = orm.Model(new(models.OwnerUnit)).Where("id = ?", ownerParams.ID)

	if err := orm.Association("AnimalUnits").Find(&ownersAnimals); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusCreated, &ownersAnimals)
}
