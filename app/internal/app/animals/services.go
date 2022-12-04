package animals

import (
	"github.com/steelWinds/animals-rest-go/internal/app/models"
	"github.com/steelWinds/animals-rest-go/internal/pkg/services"
	"gorm.io/gorm"
)

type AnimalsSet struct {
	*gorm.DB
	services.Service[models.AnimalUnit]
}

func NewService(db *gorm.DB) *AnimalsSet {
	return &AnimalsSet{
		db, 
		services.Service[models.AnimalUnit]{
			DB: db,
			Preload: true,
		},
	}
} 