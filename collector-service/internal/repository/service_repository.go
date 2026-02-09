package repository

import (
	"context"

	"github.com/Chintukr2004/collector/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ServiceRepository struct {
	db *pgxpool.Pool
}

func NewServiceRepository(db *pgxpool.Pool) *ServiceRepository {
	return &ServiceRepository{db: db}
}

func (r *ServiceRepository) GetAllServices(ctx context.Context) ([]models.Service, error) {

	query := `SELECT id, name, url, check_interval, timeout, created_at FROM services`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []models.Service

	for rows.Next() {
		var s models.Service

		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.URL,
			&s.CheckInterval,
			&s.Timeout,
			&s.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		services = append(services, s)
	}

	return services, nil
}
