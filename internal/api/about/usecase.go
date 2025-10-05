package about

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type UseCase interface {
	Update(ctx context.Context, req entity.About) (*entity.About, error)
	Get(ctx context.Context) (*entity.About, error)
}
