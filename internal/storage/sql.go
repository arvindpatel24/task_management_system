package storage

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
