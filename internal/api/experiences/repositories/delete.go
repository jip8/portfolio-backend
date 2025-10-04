package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jmoiron/sqlx"
)

type DeleteRepository struct {
	config      *entity.Config
	redisClient *redis.Client
	db          *sqlx.DB
}

func NewDeleteRepository(config *entity.Config, redisClient *redis.Client, db *sqlx.DB) *DeleteRepository {
	return &DeleteRepository{
		config:      config,
		redisClient: redisClient,
		db:          db,
	}
}

func (r *DeleteRepository) Execute(ctx context.Context, id int) error {
	query := `DELETE FROM portfolio.experiences WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}