package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var conn *pgxpool.Pool = nil

func GetConnection() *pgxpool.Pool {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is required")
	}

	if conn != nil {
		return conn
	}

	dbConnection, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	conn = dbConnection

	return conn
}
