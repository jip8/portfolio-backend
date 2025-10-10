package attachments

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type Repository interface {
	Insert(ctx context.Context, input []entity.AttachmentFlat) error
	Delete(ctx context.Context, module string, parent_id int, ids []int) error
	DeleteAll(ctx context.Context, module string, parent_id int) error
	GetListById(ctx context.Context, module string, parent_id int) (entity.AttachmentRespArray, error)
}