package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/links"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type DeleteAllUC struct {
	config         *entity.Config
	linksRepo      links.Repository
	postgresClient *services.PostgresClient
}

func NewDeleteAllUC(config *entity.Config,linksRepo links.Repository, postgresClient *services.PostgresClient) *DeleteAllUC {
	return &DeleteAllUC{
		config:         config,
		linksRepo:      linksRepo,
		postgresClient: postgresClient,
	}
}

func (u *DeleteAllUC) Execute(ctx context.Context, module string, parent_id int) error {

	err := u.linksRepo.DeleteAll(ctx, module, parent_id)
	if err != nil {
		return err
	}

	return nil
}
