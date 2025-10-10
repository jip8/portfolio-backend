package attachments

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type UseCase interface {
	Insert(ctx context.Context, input []entity.AttachmentFlat) error
	Delete(ctx context.Context, module string, parent_id int, ids []int) error
	GetListById(ctx context.Context, module string, parent_id int) (entity.AttachmentRespArray, error)
	DeleteAll(ctx context.Context, module string, parent_id int) error
	Get(ctx context.Context, objectKey string) (*entity.File, error)
}