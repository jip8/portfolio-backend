package experiences

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type UseCase interface {
	Create(ctx context.Context, req entity.ExperienceFlat) (*entity.ExperienceResp, error)
	Update(ctx context.Context, req entity.ExperienceFlat, id string) (*entity.ExperienceResp, error)
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*entity.ExperienceResp, error)
	GetList(ctx context.Context) (*entity.List, error)
}
