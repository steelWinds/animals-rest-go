package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDB(dbname, sslmode string) {
	var DB_USER = os.Getenv("POSTGRES_USER")
	var DB_PASS = os.Getenv("POSTGRES_PASSWORD")

	var DB_HOST = os.Getenv("DB_HOST")
	var DB_PORT = os.Getenv("DB_PORT")
	
	dsn := fmt.Sprintf(
		"user=%v password=%v dbname=%v sslmode=%v host=%v port=%v",
		DB_USER, DB_PASS, dbname, sslmode, DB_HOST, DB_PORT,
	)
	
	db, err := gorm.Open(postgres.Open(dsn), new(gorm.Config))

	if err != nil {
		log.Fatal(err.Error())
	}

	DB = db
}