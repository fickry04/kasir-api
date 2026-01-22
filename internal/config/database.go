package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	//Env for development from .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}

	//Get env
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		fmt.Println("DATABASE_URL is not set")
	}

	connStr := dbURL
	return sql.Open("postgres", connStr)
}
