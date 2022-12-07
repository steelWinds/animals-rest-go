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
	
	if err = service.DB.Create(&item).Error; err != nil {
		return
	}

	return
}

func (service *Service[T]) GetItem(id uint) (item T, err error) {
	db := *service.DB.Model(new(T))
	
	if service.Preload {
		db = *db.Preload(clause.Associations)
	}

	if err = db.Where("id = ?", id).First(&item).Error; err != nil {
		return
	}

	return
}

func (service *Service[T]) GetAllItems() (items []T, err error) {	
	db := *service.DB
	
	if service.Preload {
		db = *db.Preload(clause.Associations)
	}

	if err = db.Find(&items).Error; err != nil {
		return
	}

	return
}

func (service *Service[T]) DeleteItem(id uint) (err error) {
	var item T
	
	item, err = service.GetItem(id)

	if err != nil {
		return
	}

	if err = service.DB.Unscoped().Delete(&item).Error; err != nil {
		return
	}

	return
}