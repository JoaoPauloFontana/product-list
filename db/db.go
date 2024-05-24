package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	sslmode := os.Getenv("DB_SSLMODE")

	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", user, dbname, password, host, sslmode)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(fmt.Sprintf("Error opening database: %v", err))
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		panic(fmt.Sprintf("Error connecting to the database: %v", err))
	}

	return db
}
