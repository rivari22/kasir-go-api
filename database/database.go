package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver with pgx
)

func InitDB(connectionString string) (*sql.DB, error) {
	log.Println("DB Params", connectionString)
	db, err := sql.Open("pgx", connectionString)
	log.Println("DB URL:", os.Getenv("DATABASE_CONNECTION"))

	if err != nil {
		log.Println("Database error on sql open", err)
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	if err = db.Ping(); err != nil {
		log.Println("Database error on ping", err)
		return nil, err
	}

	// Set connection pool settings (optional tapi recommended)

	log.Println("Database connected successfully")
	return db, nil
}
