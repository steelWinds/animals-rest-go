package db

import (
	"github.com/jmoiron/sqlx"

	"fmt"

	"log"

	_ "github.com/lib/pq"
)



func ConnectDB(user, password, dbname, sslmode string, port int) (*sqlx.DB, error) {
	connectionStr := fmt.Sprintf(
		"user=%v password=%v dbname=%v sslmode=%v host=postgres_db",
		user, password, dbname, sslmode,
	)
	
	db, err := sqlx.Connect("postgres", connectionStr)

	if err != nil {
		log.Fatalln("Error connecting to postgres DB")
	}

	return db, nil
}