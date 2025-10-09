package links

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type UseCase interface {
	Upsert(ctx context.Context, input entity.LinkArray) error
	GetListById(ctx context.Context, module string, parent_id int) (entity.LinkRespArray, error)
	DeleteAll(ctx context.Context, module string, parent_id int) error
}
