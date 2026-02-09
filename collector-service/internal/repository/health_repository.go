package repository

import (
	"context"

	"github.com/Chintukr2004/collector/internal/checker"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthRepository struct {
	db *pgxpool.Pool
}

func NewHealthRepository(db *pgxpool.Pool) *HealthRepository {

	return &HealthRepository{db: db}
}

func (r *HealthRepository) SaveResult(ctx context.Context, result checker.Result) error {
	query := `
	INSERT INTO health_checks (service_id, status, latency_ms)
	VALUES ($1, $2, $3)
	`
	_, err := r.db.Exec(
		ctx,
		query,
		result.ServiceID,
		result.Status,
		result.Latency.Milliseconds(),
	)

	return err
}
