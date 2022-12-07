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
	Name     string `fake:"{firstname}" json:"name"`
	Age      int    `fake:"{number:1,100}" json:"age"`
	Type     string `fake:"{animaltype}" json:"type"`
	OwnerUnits []OwnerUnit `fake:"skip" json:"ownerUnits" gorm:"many2many:own_maps;association_jointable_foreignkey:owner_unit_id;save_association:false"`
	Gorm `fake:"skip"`
}

type OwnMap struct {
	AnimalUnitID uint `json:"animalId"`
	OwnerUnitID uint `json:"ownerId"`
	Gorm
}