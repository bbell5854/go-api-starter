package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// DB : Exported database connection
var DB *sqlx.DB

// Connect : Creates a database connection with the provded connection string
func Connect(databaseType string, connString string) error {
	db, err := sqlx.Connect(databaseType, connString)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	DB = db
	return nil
}
