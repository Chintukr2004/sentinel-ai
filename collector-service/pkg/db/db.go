package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB() *pgxpool.Pool {

	dsn := "postgres://admin:admin@postgres:5432/sentinel"

	var dbpool *pgxpool.Pool
	var err error

	for i := 0; i < 10; i++ {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		dbpool, err = pgxpool.New(ctx, dsn)
		if err == nil {

			err = dbpool.Ping(ctx)
			if err == nil {
				cancel()
				log.Println("Connected to PostgreSQL")
				return dbpool
			}
		}

		cancel()

		log.Println("⏳ Database not ready... retrying in 2 seconds")
		time.Sleep(2 * time.Second)
	}

	log.Fatal("Could not connect to database after retries")

	return nil
}
