package repositories

import (
	"context"
	"time"

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
	initialDate, err := time.Parse("2006-01-02", *req.InitialDate)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02", *req.EndDate)
	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO experiences (title, function, description, initial_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var newID int
	err = r.db.QueryRowContext(ctx, query, req.Title, *req.Function, *req.Description, initialDate, endDate).Scan(&newID)
	if err != nil {
		return nil, err
	}

	return &newID, nil
}