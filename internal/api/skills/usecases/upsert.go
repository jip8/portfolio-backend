package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/skills"
	"github.com/jip/portfolio-backend/internal/entity"

	"github.com/jip/portfolio-backend/internal/services"
)

type UpsertUC struct {
	config         *entity.Config
	skillsRepo     skills.Repository
	postgresClient *services.PostgresClient
}

func NewUpsertUC(config *entity.Config, skillsRepo skills.Repository, postgresClient *services.PostgresClient) *UpsertUC {
	return &UpsertUC{
		config:         config,
		skillsRepo:     skillsRepo,
		postgresClient: postgresClient,
	}
}

func (u *UpsertUC) Execute(ctx context.Context, parent_id *int, module *string, input entity.SkillArray) error {
	var err error

	err = input.Validate()
	if err != nil {
		return err
	}

	ctx, err = u.postgresClient.StartProcess(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = u.postgresClient.CloseProcess(ctx, err)
	}()

	var toDelete []int
	var toInsert []int
	var toUpsert []entity.SkillFlat

	for _, skill := range input {
		if skill.Id == nil || *skill.Id == 0 {
			toUpsert = append(toUpsert, skill)
			continue
		} else if *skill.Id < 0 {
			toDelete = append(toDelete, -*skill.Id)
			continue
		} else {
			toInsert = append(toInsert, *skill.Id)
			continue
		}
	}

	err = u.skillsRepo.Delete(ctx, toDelete)
	if err != nil {
		return err
	}

	newIds, err := u.skillsRepo.Upsert(ctx, toUpsert)
	if err != nil {
		return err
	}

	toInsert = append(toInsert, newIds...)

	if parent_id != nil && module != nil {
		err = u.skillsRepo.AddExclusive(ctx, parent_id, module, toInsert)
		if err != nil {
			return err
		}
	}

	return nil
}
