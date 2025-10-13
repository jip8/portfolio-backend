package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/attachments"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetListByIdUC struct {
	config          *entity.Config
	attachmentsRepo attachments.Repository
	postgresClient  *services.PostgresClient
}

func NewGetListByIdUC(config *entity.Config, attachmentsRepo attachments.Repository, postgresClient *services.PostgresClient) *GetListByIdUC {
	return &GetListByIdUC{
		config:          config,
		attachmentsRepo: attachmentsRepo,
		postgresClient:  postgresClient,
	}
}

func (u *GetListByIdUC) Execute(ctx context.Context, module string, parent_id int) (entity.AttachmentRespArray, error) {

	resp, err := u.attachmentsRepo.GetListById(ctx, module, parent_id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
