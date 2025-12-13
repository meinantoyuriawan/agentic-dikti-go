package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

var (
	database = os.Getenv("POSTGRES_DB")
	password = os.Getenv("POSTGRES_PASSWORD")
	username = os.Getenv("POSTGRES_USER")
	port     = os.Getenv("POSTGRES_PORT")
	host     = os.Getenv("POSTGRES_HOST")
	schema   = os.Getenv("POSTGRES_SCHEMA")
)

func New() *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
