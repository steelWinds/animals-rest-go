package resources

import (
	"log"

	"gorm.io/gorm"
)

func Clear(db *gorm.DB) {
	var tables []string
	var err error

	if tables, err = db.Migrator().GetTables(); err != nil {
		log.Fatal(err)
	}

	for _, table := range tables {
		db.Migrator().DropTable(table)
	}
}
