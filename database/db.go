package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/joho/godotenv"
	// _ "github.com/tursodatabase/go-libsql"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func Connect() *sql.DB {
	db, err := sql.Open("libsql", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	return db
}
