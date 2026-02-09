package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB() *pgxpool.Pool {

	dsn := "postgres://admin:admin@localhost:5433/sentinel"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	err = dbpool.Ping(ctx)
	if err != nil {
		log.Fatalf("Database not reachable: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL")

	return dbpool
}
