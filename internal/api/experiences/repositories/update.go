package repositories

import (
	"context"
	"errors"

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
	query := `
		UPDATE portfolio.experiences
		SET
			title = :title,
			"function" = :function,
			description = :description,
			initial_date = :initial_date_time,
			end_date = :end_date_time,
			actual = :actual
		WHERE
			id = :id
	`

	namedQuery, args, err := r.db.BindNamed(query, req)
	if err != nil {
		return nil, err
	}

	result, err := r.db.ExecContext(ctx, namedQuery, args...)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("experience not found")
	}

	return req.Id, nil
}