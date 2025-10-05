package articles

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type UseCase interface {
	Create(ctx context.Context, req entity.ArticleFlat) (*entity.ArticleResp, error)
	Update(ctx context.Context, req entity.ArticleFlat) (*entity.ArticleResp, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*entity.ArticleResp, error)
	GetList(ctx context.Context, listReq entity.ListReq) (*entity.List[entity.ArticleResp], error)
}
