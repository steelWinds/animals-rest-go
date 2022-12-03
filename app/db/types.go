package db

import (
	"time"

	"gorm.io/gorm"
)

type IDParams struct {
	ID uint `uri:"id" binding:"required"`
}

type Gorm struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time		 `json:"createdAt"`
	UpdatedAt time.Time		 `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type AnimalUnit struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Type     string `json:"type"`
	OwnerUnits []OwnerUnit `gorm:"many2many:own_maps;association_jointable_foreignkey:owner_unit_id;save_association:false"`
	Gorm
}

type OwnerUnit struct {
	Name    string `json:"name"`
	AnimalUnits []AnimalUnit `gorm:"many2many:own_maps;association_jointable_foreignkey:animal_unit_id;save_association:false"`
	Gorm
}

type OwnMap struct {
	AnimalUnitID uint `json:"animalId"`
	OwnerUnitID uint `json:"ownerId"`
	Gorm
}