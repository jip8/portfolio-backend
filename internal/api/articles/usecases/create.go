package usecases

import (
	"fmt"
	"context"

	"github.com/jip/portfolio-backend/internal/api/articles"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jip/portfolio-backend/internal/api/links"
	"github.com/jip/portfolio-backend/internal/api/skills"
)

type CreateUC struct {
	config         *entity.Config
	articlesRepo   articles.Repository
	byId           *GetByIdUC
	postgresClient *services.PostgresClient
	linksUC        links.UseCase
	skillsUC       skills.UseCase
}

func NewCreateUC(config *entity.Config, articlesRepo articles.Repository, byId *GetByIdUC, postgresClient *services.PostgresClient, linksUC links.UseCase, skillsUC skills.UseCase) *CreateUC {
	return &CreateUC{
		config:         config,
		articlesRepo:   articlesRepo,
		byId:           byId,
		postgresClient: postgresClient,
		linksUC:        linksUC,
		skillsUC:       skillsUC,
	}
}

func (u *CreateUC) Execute(ctx context.Context, req entity.ArticleFlat) (resp *entity.ArticleResp, err error) {
	err = req.Validate()
	if err != nil {
		return nil, err
	}

	ctx, err = u.postgresClient.StartProcess(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = u.postgresClient.CloseProcess(ctx, err)
	}()

	var createdId *int
	createdId, err = u.articlesRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	resp, err = u.byId.Execute(ctx, *createdId)
	if err != nil {
		return nil, err
	}

	module := fmt.Sprintf("%s", ModuleName)
	err = u.skillsUC.Upsert(ctx, createdId, &module, req.Skills)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
