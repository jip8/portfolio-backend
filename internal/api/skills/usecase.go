package skills

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type UseCase interface {
	Upsert(ctx context.Context, parent_id *int, module *string, input entity.SkillArray) error
	GetList(ctx context.Context, module *string, parent_id *int) (entity.SkillRespArray, error)
	DeleteAll(ctx context.Context, module string, parent_id int) error
}
