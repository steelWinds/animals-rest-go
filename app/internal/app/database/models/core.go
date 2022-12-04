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


type IDParams struct {
	ID uint `uri:"id" binding:"required"`
}