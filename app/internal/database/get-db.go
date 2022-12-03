package database

import "gorm.io/gorm"

func GetDB() *gorm.DB {
	newDB := *DB
	
	return &newDB
}