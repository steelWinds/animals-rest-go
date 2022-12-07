package database

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
	"github.com/tanimutomo/sqlfile"
)

func ConnectDB() (db *gorm.DB, err error) {	
	var (
		DB_USER = os.Getenv("POSTGRES_USER")
		DB_PASS = os.Getenv("POSTGRES_PASSWORD")
		DB_NAME = os.Getenv("DB_NAME")
		DB_HOST = os.Getenv("DB_HOST")
		DB_PORT = os.Getenv("DB_PORT")
		SSL_MODE = os.Getenv("SSL_MODE")
	)
	
	dsn := fmt.Sprintf(
		"user=%v password=%v dbname=%v sslmode=%v host=%v port=%v",
		DB_USER, DB_PASS, DB_NAME, SSL_MODE, DB_HOST, DB_PORT,
	)
	
	db, err = gorm.Open(postgres.Open(dsn), new(gorm.Config))

	if err != nil {
		return
	}

	return
}

func SetTestMigrations(db *sql.DB) (err error) {
	sqlExec := sqlfile.New()

	if err = sqlExec.File("../configurate.sql"); err != nil {
		return
	}

	sqlExec.Exec(db)

	return
}