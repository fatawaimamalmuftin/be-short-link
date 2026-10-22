package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB() *pgxpool.Pool {
	dbHost := os.Getenv("DF_DB_HOST")
	dbPort := os.Getenv("DF_DB_PORT_H")
	dbUser := os.Getenv("DF_DB_USER")
	dbPass := os.Getenv("DF_DB_PASS")
	dbName := os.Getenv("DF_DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Println("Unable to connect to database: ", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Println("Database ping failed: ", err)
	}

	log.Println("Connected to PostgreSQL successfully")
	return pool
}
