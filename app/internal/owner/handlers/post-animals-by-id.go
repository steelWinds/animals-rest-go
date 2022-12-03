package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelWinds/animals-rest-go/db"
	"github.com/steelWinds/animals-rest-go/internal/database"
	"github.com/steelWinds/animals-rest-go/internal/models"
)

func PostAnimlsById(c *gin.Context) {
	var ownerParams db.IDParams

	if err := c.ShouldBindUri(&ownerParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })

		return
	}

	var animals []db.AnimalUnit

	if err := c.BindJSON(&animals); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })

		return
	}

	var ownMappers []db.OwnMap
	
	for _, animal := range animals {
		ownMappers = append(ownMappers, db.OwnMap{AnimalUnitID: animal.ID, OwnerUnitID: ownerParams.ID})
	}

	orm := database.GetDB()

	if tx := orm.Create(&ownMappers); tx.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"message": tx.Error.Error()})

		return
	}

	var ownersAnimals []db.AnimalUnit

	orm = orm.Model(new(models.OwnerUnit)).Where("id = ?", ownerParams.ID)

	if err := orm.Association("AnimalUnits").Find(&ownersAnimals); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusCreated, &ownersAnimals)
}
