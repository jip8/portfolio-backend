package courses

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, req entity.CourseFlat) (*int, error)
	Update(ctx context.Context, req entity.CourseFlat) (*int, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*entity.CourseResp, error)
	GetList(ctx context.Context, listReq entity.ListReq) (*entity.List[entity.CourseResp], error)
}
