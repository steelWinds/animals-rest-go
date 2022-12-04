package owners

import (
	"github.com/steelWinds/animals-rest-go/internal/app/models"
	"github.com/steelWinds/animals-rest-go/internal/pkg/services"
	"gorm.io/gorm"
)

type OwnersSet struct {
	*gorm.DB
	services.Service[models.OwnerUnit]
}

func NewService(db *gorm.DB) *OwnersSet {
	return &OwnersSet{
		db, 
		services.Service[models.OwnerUnit]{
			DB: db,
			Preload: true,
		},
	}
}

func (ctx *OwnersSet) PostAnimals(
	id uint,
	animals []models.AnimalUnit,
) (item models.OwnerUnit, err error) {
	var ownMappers []models.OwnMap
	
	for _, animal := range animals {
		ownMappers = append(
			ownMappers, 
			models.OwnMap{
				AnimalUnitID: animal.ID,
				OwnerUnitID: id,
			},
		)
	}

	if tx := ctx.DB.Create(&ownMappers); tx.Error != nil {
		err = tx.Error

		return
	}

	item, err = ctx.GetItem(id)

	if err != nil {
		return
	}

	return
}