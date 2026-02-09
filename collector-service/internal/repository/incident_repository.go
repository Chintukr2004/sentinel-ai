package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IncidentRepository struct {
	db *pgxpool.Pool
}

func NewIncidentRepository(db *pgxpool.Pool) *IncidentRepository {
	return &IncidentRepository{db: db}
}

// Check if an active incident exists
func (r *IncidentRepository) HasActiveIncident(ctx context.Context, serviceID int) (bool, error) {

	query := `
	SELECT EXISTS (
		SELECT 1 FROM incidents
		WHERE service_id=$1 AND resolved_at IS NULL
	)`

	var exists bool

	err := r.db.QueryRow(ctx, query, serviceID).Scan(&exists)

	return exists, err
}

// Create incident
func (r *IncidentRepository) CreateIncident(ctx context.Context, serviceID int) error {

	query := `
	INSERT INTO incidents (service_id, status)
	VALUES ($1, 'ONGOING')
	`

	_, err := r.db.Exec(ctx, query, serviceID)

	return err
}

// Resolve incident
func (r *IncidentRepository) ResolveIncident(ctx context.Context, serviceID int) error {

	query := `
	UPDATE incidents
	SET resolved_at = CURRENT_TIMESTAMP,
	    status = 'RESOLVED'
	WHERE service_id=$1
	AND resolved_at IS NULL
	`

	_, err := r.db.Exec(ctx, query, serviceID)

	return err
}
