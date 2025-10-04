package repositories

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GetByIdRepository struct {
	config      *entity.Config
	redisClient *redis.Client
	db          *sqlx.DB
}

func NewGetByIdRepository(config *entity.Config, redisClient *redis.Client, db *sqlx.DB) *GetByIdRepository {
	return &GetByIdRepository{
		config:      config,
		redisClient: redisClient,
		db:          db,
	}
}

func (r *GetByIdRepository) Execute(ctx context.Context, id int) (*entity.ExperienceResp, error) {
	query := `SELECT * FROM experiences WHERE id = $1`
	var experience entity.Experience
	if err := r.db.GetContext(ctx, &experience, query, id); err != nil {
		return nil, err
	}

	initialDate := experience.InitialDate.Format("2006-01-02")
	endDate := experience.EndDate.Format("2006-01-02")

	return &entity.ExperienceResp{
		Id:          strconv.Itoa(experience.ID),
		Title:       experience.Title,
		Function:    &experience.Function,
		Description: &experience.Description,
		InitialDate: &initialDate,
		EndDate:     &endDate,
	}, nil
}