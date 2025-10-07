package links

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type Repository interface {
	Upsert(ctx context.Context, input entity.LinkArray) error
	Delete(ctx context.Context, ids []int) error
	DeleteAll(ctx context.Context, module string, parent_id int) error
	GetListById(ctx context.Context, module string, parent_id int) (entity.LinkRespArray, error)
}