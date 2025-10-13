package usecases

import (
	"fmt"
	"context"

	"github.com/jip/portfolio-backend/internal/api/articles"
	"github.com/jip/portfolio-backend/internal/entity"

	"github.com/jip/portfolio-backend"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jip/portfolio-backend/internal/api/links"
	"github.com/jip/portfolio-backend/internal/api/skills"
)

type UpdateUC struct {
	config         *entity.Config
	articlesRepo   articles.Repository
	byId           *GetByIdUC
	postgresClient *services.PostgresClient
	linksUC        links.UseCase
	skillsUC       skills.UseCase
}

func NewUpdateUC(config *entity.Config, articlesRepo articles.Repository, byId *GetByIdUC, postgresClient *services.PostgresClient, linksUC links.UseCase, skillsUC skills.UseCase) *UpdateUC {
	return &UpdateUC{
		config:         config,
		articlesRepo:   articlesRepo,
		byId:           byId,
		postgresClient: postgresClient,
		linksUC:        linksUC,
		skillsUC:       skillsUC,
	}
}

func (u *UpdateUC) Execute(ctx context.Context, req entity.ArticleFlat) (resp *entity.ArticleResp, err error) {
	if req.Id == nil {
		return nil, portfolio.ErrArticleIdIsRequired
	}

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

	var updatedId *int
	updatedId, err = u.articlesRepo.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	module := fmt.Sprintf("%s", ModuleName)

	for i := range req.LinksArray {
		req.LinksArray[i].ParentId = updatedId
		req.LinksArray[i].Module = &module
	}

	err = u.linksUC.Upsert(ctx, req.LinksArray)
	if err != nil {
		return nil, err
	}

	err = u.skillsUC.Upsert(ctx, updatedId, &module, req.Skills)
	if err != nil {
		return nil, err
	}

	resp, err = u.byId.Execute(ctx, *updatedId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
