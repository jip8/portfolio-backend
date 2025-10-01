package login

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
)

type UseCase interface {
	Login(ctx context.Context, req entity.LoginRequest, ip string) (string, error)
}
