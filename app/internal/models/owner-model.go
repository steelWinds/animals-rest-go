package models

import (
	"github.com/steelWinds/animals-rest-go/internal/database/models"
)

type OwnerUnit struct {
	Name    string `json:"name"`
	AnimalUnits []AnimalUnit `json:"animalUnits" gorm:"many2many:own_maps;association_jointable_foreignkey:animal_unit_id;save_association:false"`
	models.Gorm
}