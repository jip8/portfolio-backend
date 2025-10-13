package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/links"
	"github.com/jip/portfolio-backend/internal/entity"

	"github.com/jip/portfolio-backend/internal/services"
)

type UpsertUC struct {
	config         *entity.Config
	linksRepo      links.Repository
	postgresClient *services.PostgresClient
}

func NewUpsertUC(config *entity.Config, linksRepo links.Repository, postgresClient *services.PostgresClient) *UpsertUC {
	return &UpsertUC{
		config:         config,
		linksRepo:      linksRepo,
		postgresClient: postgresClient,
	}
}

func (u *UpsertUC) Execute(ctx context.Context, input entity.LinkArray) error {
	var err error
	
	err = input.Validate()
	if err != nil {
		return err
	}

	ctx, err = u.postgresClient.StartProcess(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = u.postgresClient.CloseProcess(ctx, err)
	}()

	var toDelete []int
	var toUpsert []entity.LinkFlat

	for _, link := range input {
		if link.Id == nil {
			continue
		}

		id := *link.Id
		switch {
		case id >= 0:
			toUpsert = append(toUpsert, link)
		case id < 0:
			toDelete = append(toDelete, -id)
		}
	}

	if len(toDelete) > 0 {
		err = u.linksRepo.Delete(ctx, toDelete)
		if err != nil {
			return err
		}
	}

	if len(toUpsert) > 0 {
		err = u.linksRepo.Upsert(ctx, toUpsert)
		if err != nil {
			return err
		}
	}

	return nil
}
