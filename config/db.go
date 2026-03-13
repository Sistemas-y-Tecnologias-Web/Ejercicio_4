package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func NewDB() *pgxpool.Pool {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env vars")
	}

	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=require",
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASS", "postgres"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_NAME", "gamecenter"),
	)

	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatal("Error connecting to DB: ", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("DB not reachable: ", err)
	}

	log.Println("✅ connected to Postgres")
	return pool
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
