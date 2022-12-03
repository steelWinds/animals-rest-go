package models

import (
	"github.com/steelWinds/animals-rest-go/internal/database/models"
)

type OwnMap struct {
	AnimalUnitID uint `json:"animalId"`
	OwnerUnitID uint `json:"ownerId"`
	models.Gorm
}