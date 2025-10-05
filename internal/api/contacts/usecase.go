package contacts

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type UseCase interface {
	Create(ctx context.Context, req entity.Contact) (*entity.Contact, error)
	Update(ctx context.Context, req entity.Contact) (*entity.Contact, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*entity.Contact, error)
	GetList(ctx context.Context, listReq entity.ListReq) (*entity.List[entity.Contact], error)
}
