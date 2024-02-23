package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func Connect() *sql.DB {
	db, err := sql.Open("libsql", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	res, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err, res)
	}

	return db
}
