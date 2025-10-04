package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/api/experiences"

	"github.com/jip/portfolio-backend"
)

type UpdateUC struct {
	config      	*entity.Config
	redisClient 	*redis.Client
	experiencesRepo experiences.Repository
	byId 			*GetByIdUC
}

func NewUpdateUC(config *entity.Config, redisClient *redis.Client, experiencesRepo experiences.Repository, byId *GetByIdUC) *UpdateUC {
	return &UpdateUC{
		config:      		config,
		redisClient: 		redisClient,
		experiencesRepo: 	experiencesRepo,
		byId: 				byId,
	}
}

func (u *UpdateUC) Execute(ctx context.Context, req entity.ExperienceFlat) (*entity.ExperienceResp, error) {
	
	var updatedId *int

	if req.Id == nil {
		return nil, portfolio.ErrExperienceIdIsRequired
	}

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	updatedId, err = u.experiencesRepo.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	resp, err := u.byId.Execute(ctx, *updatedId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}