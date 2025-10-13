package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/skills"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetListByIdUC struct {
	config         *entity.Config
	skillsRepo     skills.Repository
	postgresClient *services.PostgresClient
}

func NewGetListByIdUC(config *entity.Config, skillsRepo skills.Repository, postgresClient *services.PostgresClient) *GetListByIdUC {
	return &GetListByIdUC{
		config:         config,
		skillsRepo:     skillsRepo,
		postgresClient: postgresClient,
	}
}

func (u *GetListByIdUC) Execute(ctx context.Context, module *string, parent_id *int) (entity.SkillRespArray, error) {

	resp, err := u.skillsRepo.GetListById(ctx, module, parent_id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
