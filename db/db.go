// Set Up PostgreSQL:
// Create the messages table in your database.
// Configure environment variables for the PostgreSQL connection string.
package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading the .env file")
	}

	var errDB error
	Pool, errDB = pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if errDB != nil {
		log.Fatalf("Unable to connect to database %v", errDB)
	}
}
