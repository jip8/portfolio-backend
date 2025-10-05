package articles

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, req entity.ArticleFlat) (*int, error)
	Update(ctx context.Context, req entity.ArticleFlat) (*int, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*entity.ArticleResp, error)
	GetList(ctx context.Context, listReq entity.ListReq) (*entity.List[entity.ArticleResp], error)
}
