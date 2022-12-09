package tests

import (
	"database/sql"
	"log"

	"github.com/tanimutomo/sqlfile"
	"gorm.io/gorm"
)

func SetTestMigrations(db *sql.DB) (err error) {
	sqlExec := sqlfile.New()

	if err = sqlExec.File("resources/configurate.sql"); err != nil {
		return
	}

	sqlExec.Exec(db)

	return
}

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
