package factory

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Conn -
func Conn() *sql.DB {
	db, err := sql.Open("postgres", "")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	log.Println("Database init.")
	return db
}
