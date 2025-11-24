package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetByIdRepository struct {
	config         *entity.Config
	postgresClient *services.PostgresClient
}

func NewGetByIdRepository(config *entity.Config,  postgresClient *services.PostgresClient) *GetByIdRepository {
	return &GetByIdRepository{
		config:         config,
		postgresClient: postgresClient,
	}
}

func (r *GetByIdRepository) Execute(ctx context.Context, id int) (*entity.ArticleResp, error) {
	query := `
	SELECT
		id,
		type,
		title,
		description,
		local,
		published_at AS published_at_time,
		revelance,
		thumbnail_id
	FROM portfolio.articles WHERE id = $1`

	var article entity.ArticleResp

	executor := r.postgresClient.GetExecutor(ctx)
	if err := executor.GetContext(ctx, &article, query, id); err != nil {
		return nil, err
	}

	return &article, nil
}