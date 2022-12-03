package database

import (
	"fmt"
	"os"
	"time"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDB(user, password, dbname, sslmode string, port int) error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              time.Second,   // Slow SQL threshold
			LogLevel:                   logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,          // Disable color
		},
	)
	
	dsn := fmt.Sprintf(
		"user=%v password=%v dbname=%v sslmode=%v host=postgres_db",
		user, password, dbname, sslmode,
	)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return err
	}

	DB = db

	return nil
}