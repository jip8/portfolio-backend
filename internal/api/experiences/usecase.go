package experiences

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type UseCase interface {
	Create(ctx context.Context, req entity.ExperienceFlat) (*entity.ExperienceResp, error)
	Update(ctx context.Context, req entity.ExperienceFlat) (*entity.ExperienceResp, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*entity.ExperienceResp, error)
	GetList(ctx context.Context, listReq entity.ListReq) (*entity.List[entity.ExperienceResp], error)
}
