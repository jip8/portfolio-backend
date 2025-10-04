package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CreateRepository struct {
	config      *entity.Config
	redisClient *redis.Client
	db          *sqlx.DB
}

func NewCreateRepository(config *entity.Config, redisClient *redis.Client, db *sqlx.DB) *CreateRepository {
	return &CreateRepository{
		config:      config,
		redisClient: redisClient,
		db:          db,
	}
}

func (r *CreateRepository) Execute(ctx context.Context, req entity.ExperienceFlat) (*int, error) {
	query := `
		INSERT INTO experiences (title, "function", description, initial_date, end_date)
		VALUES (:title, :function, :description, :initial_date, :end_date)
		RETURNING id
	`

	namedQuery, args, err := r.db.BindNamed(query, req)
	if err != nil {
		return nil, err
	}

	var newID int
	err = r.db.QueryRowContext(ctx, namedQuery, args...).Scan(&newID)
	if err != nil {
		return nil, err
	}

	return &newID, nil
}