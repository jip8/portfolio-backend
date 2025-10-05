package projects

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type UseCase interface {
	Create(ctx context.Context, req entity.ProjectFlat) (*entity.ProjectResp, error)
	Update(ctx context.Context, req entity.ProjectFlat) (*entity.ProjectResp, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*entity.ProjectResp, error)
	GetList(ctx context.Context, listReq entity.ListReq) (*entity.List[entity.ProjectResp], error)
}
