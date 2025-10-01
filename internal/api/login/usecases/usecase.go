package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/login"
	"github.com/jip/portfolio-backend/internal/entity"
)

type loginUC struct {
	login *LoginUC
}

func NewLoginUseCase(config *entity.Config, redisClient *redis.Client) login.UseCase {
	return &loginUC{
		login: NewLoginUC(config, redisClient),
	}
}

func (u *loginUC) Login(ctx context.Context, req entity.LoginRequest, ip string) (string, error) {
	return u.login.Execute(ctx, req, ip)
}
