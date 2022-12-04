package models

import (
	"time"

	"gorm.io/gorm"
)

type Gorm struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time		 `json:"createdAt"`
	UpdatedAt time.Time		 `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type OwnerUnit struct {
	Name    string `json:"name"`
	AnimalUnits []AnimalUnit `json:"animalUnits" gorm:"many2many:own_maps;association_jointable_foreignkey:animal_unit_id;save_association:false"`
	Gorm
}

type AnimalUnit struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Type     string `json:"type"`
	OwnerUnits []OwnerUnit `json:"ownerUnits" gorm:"many2many:own_maps;association_jointable_foreignkey:owner_unit_id;save_association:false"`
	Gorm
}

type OwnMap struct {
	AnimalUnitID uint `json:"animalId"`
	OwnerUnitID uint `json:"ownerId"`
	Gorm
}