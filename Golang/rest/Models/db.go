package models

import (
	"database/sql"
	"log"
)

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "feelme_admin:zj439$Z2p@tcp(119.59.96.90)/feelme_db")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	return db
}
