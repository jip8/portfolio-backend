package repositories

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UpdateRepository struct {
	config      *entity.Config
	redisClient *redis.Client
	db          *sqlx.DB
}

func NewUpdateRepository(config *entity.Config, redisClient *redis.Client, db *sqlx.DB) *UpdateRepository {
	return &UpdateRepository{
		config:      config,
		redisClient: redisClient,
		db:          db,
	}
}

func (r *UpdateRepository) Execute(ctx context.Context, req entity.ExperienceFlat) (*int, error) {
	initialDate, err := time.Parse("2006-01-02", *req.InitialDate)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02", *req.EndDate)
	if err != nil {
		return nil, err
	}

	query := `
		UPDATE experiences
		SET title = $1, function = $2, description = $3, initial_date = $4, end_date = $5
		WHERE id = $6
	`

	_, err = r.db.ExecContext(ctx, query, req.Title, *req.Function, *req.Description, initialDate, endDate, *req.Id)
	if err != nil {
		return nil, err
	}

	return req.Id, nil
}