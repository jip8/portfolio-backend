package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/courses"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type DeleteUC struct {
	config         *entity.Config
	coursesRepo    courses.Repository
	postgresClient *services.PostgresClient
}

func NewDeleteUC(config *entity.Config, coursesRepo courses.Repository, postgresClient *services.PostgresClient) *DeleteUC {
	return &DeleteUC{
		config:         config,
		coursesRepo:    coursesRepo,
		postgresClient: postgresClient,
	}
}

func (u *DeleteUC) Execute(ctx context.Context, id int) (err error) {
	ctx, err = u.postgresClient.StartProcess(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = u.postgresClient.CloseProcess(ctx, err)
	}()

	err = u.coursesRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
