package services

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Service[T any] struct {
	DB *gorm.DB
	Preload bool
}

func (service *Service[T]) CreateItem(createItem T) (item T, err error)  {
	item = createItem
	
	if tx := service.DB.Create(&item); tx.Error != nil {
		err = tx.Error
	}

	return
}

func (service *Service[T]) GetItem(id uint) (item T, err error) {
	db := *service.DB.Model(new(T))
	
	if service.Preload {
		db = *db.Preload(clause.Associations)
	}

	if tx := db.Where("id = ?", id).First(&item); tx.Error != nil {
		err = tx.Error

		return
	}

	return
}

func (service *Service[T]) GetAllItems() (items []T, err error) {	
	db := *service.DB
	
	if service.Preload {
		db = *db.Preload(clause.Associations)
	}

	if tx := db.Find(&items); tx.Error != nil {
		err = tx.Error

		return
	}

	return
}