package skills

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type Repository interface {
	Upsert(ctx context.Context, input entity.SkillArray) ([]int, error)
	Delete(ctx context.Context, ids []int) error
	DeleteAll(ctx context.Context, module string, parent_id int) error
	GetListById(ctx context.Context, module *string, parent_id *int) (entity.SkillRespArray, error)
	GetList(ctx context.Context) (entity.SkillRespArray, error)
	AddExclusive(ctx context.Context, parent_id *int, module *string, ids []int) error
}
