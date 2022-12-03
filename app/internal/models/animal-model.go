package models

import (
	"github.com/steelWinds/animals-rest-go/internal/database/models"
)

type AnimalUnit struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Type     string `json:"type"`
	OwnerUnits []OwnerUnit `gorm:"many2many:own_maps;association_jointable_foreignkey:owner_unit_id;save_association:false"`
	models.Gorm
}