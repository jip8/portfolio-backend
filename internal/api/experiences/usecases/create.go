package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/experiences"
	"github.com/jip/portfolio-backend/internal/entity"
)

type CreateUC struct {
	config      	*entity.Config
	redisClient 	*redis.Client
	experiencesRepo experiences.Repository
	byId 			*GetByIdUC
}

func NewCreateUC(config *entity.Config, redisClient *redis.Client, experiencesRepo experiences.Repository, byId *GetByIdUC) *CreateUC {
	return &CreateUC{
		config:      		config,
		redisClient: 		redisClient,
		experiencesRepo: 	experiencesRepo,
		byId: 				byId,
	}
}

func (u *CreateUC) Execute(ctx context.Context, req entity.ExperienceFlat) (*entity.ExperienceResp, error) {
	
	var createdId *int

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	createdId, err = u.experiencesRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	resp, err := u.byId.Execute(ctx, *createdId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}