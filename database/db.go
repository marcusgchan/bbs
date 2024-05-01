package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/tursodatabase/go-libsql"
)

func Connect() *sql.DB {
	fmt.Printf("db url |%v|", os.Getenv("DB_URL"))
	db, err := sql.Open("libsql", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	return db
}
