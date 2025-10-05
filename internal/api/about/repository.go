package about

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type Repository interface {
	Update(ctx context.Context, req entity.About) error
	Get(ctx context.Context) (*entity.About, error)
}
