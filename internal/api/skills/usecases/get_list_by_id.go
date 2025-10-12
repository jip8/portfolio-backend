package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/skills"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetListUC struct {
	config         *entity.Config
	skillsRepo     skills.Repository
	postgresClient *services.PostgresClient
}

func NewGetListUC(config *entity.Config, skillsRepo skills.Repository, postgresClient *services.PostgresClient) *GetListUC {
	return &GetListUC{
		config:         config,
		skillsRepo:     skillsRepo,
		postgresClient: postgresClient,
	}
}

func (u *GetListUC) Execute(ctx context.Context, module *string, parent_id *int) (entity.SkillRespArray, error) {

	resp, err := u.skillsRepo.GetListById(ctx, module, parent_id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
