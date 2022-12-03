package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelWinds/animals-rest-go/db"
	"gorm.io/gorm"
)

func PostAnimlsByIdH(orm *gorm.DB) func(*gin.Context) {
	handler := func (c *gin.Context) {
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

		if tx := orm.Create(&ownMappers); tx.Error != nil {
			c.JSON(http.StatusConflict, gin.H{"message": tx.Error.Error()})

			return
		} 

		var owner = db.OwnerUnit{
			Gorm: db.Gorm{ID: ownerParams.ID},
		}
		var ownersAnimals []db.AnimalUnit

		orm.Model(&owner).Association("AnimalUnits").Find(&ownersAnimals)

		c.JSON(http.StatusCreated, &ownersAnimals)
	}

	return handler
}
