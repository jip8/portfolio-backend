package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/skills"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type DeleteAllUC struct {
	config         *entity.Config
	skillsRepo     skills.Repository
	postgresClient *services.PostgresClient
}

func NewDeleteAllUC(config *entity.Config, skillsRepo skills.Repository, postgresClient *services.PostgresClient) *DeleteAllUC {
	return &DeleteAllUC{
		config:         config,
		skillsRepo:     skillsRepo,
		postgresClient: postgresClient,
	}
}

func (u *DeleteAllUC) Execute(ctx context.Context, module string, parent_id int) error {

	err := u.skillsRepo.DeleteAll(ctx, module, parent_id)
	if err != nil {
		return err
	}

	return nil
}
