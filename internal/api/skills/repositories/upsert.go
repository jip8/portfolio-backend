package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type UpsertRepository struct {
	config         *entity.Config
	postgresClient *services.PostgresClient
}

func NewUpsertRepository(config *entity.Config, postgresClient *services.PostgresClient) *UpsertRepository {
	return &UpsertRepository{
		config:         config,
		postgresClient: postgresClient,
	}
}

func (r *UpsertRepository) Execute(ctx context.Context, skills []entity.SkillFlat) ([]int, error) {
	if len(skills) == 0 {
		return nil, nil
	}

	valueStrings := make([]string, 0, len(skills))
	valueArgs := make([]interface{}, 0, len(skills) * 4)
	i := 1
	for _, skill := range skills {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", i, i+1, i+2, i+3))
		if *skill.Id == 0 {
			valueArgs = append(valueArgs, nil)
		} else {
			valueArgs = append(valueArgs, skill.Id)
		}

		valueArgs = append(valueArgs, skill.Title)
		valueArgs = append(valueArgs, skill.Revelance)
		valueArgs = append(valueArgs, skill.Description)
		i += 4
	}

	stmt := `
        INSERT INTO portfolio.skills (id, title, revelance, description)
        VALUES %s
        ON CONFLICT (id) DO UPDATE SET
            skill = EXCLUDED.skill,
            revelance = EXCLUDED.revelance,
            description = EXCLUDED.description,
            updated_at = NOW()
		RETURNING id
    `
	query := fmt.Sprintf(stmt, strings.Join(valueStrings, ","))

	executor := r.postgresClient.GetExecutor(ctx)
	rows, err := executor.QueryContext(ctx, query, valueArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}
