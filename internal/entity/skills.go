package entity

import "github.com/jip/portfolio-backend"

type SkillFlat struct {
	Id              *int       `json:"id" db:"id"`
	Title           *string    `json:"title" db:"title"`
	Revelance       *int       `json:"revelance" db:"revelance"`
	Description     *string    `json:"description" db:"description"`
	UpdatedAt       *string    `json:"updated_at" db:"updated_at"`
}

type SkillArray []SkillFlat

func (s SkillArray) Validate() error {
	for _, skill := range s {
		if skill.Id == nil {
			return portfolio.ErrSkillIdIsRequired
		}

		if skill.Title == nil || *skill.Title == "" {
			return portfolio.ErrSkillTitleIsRequired
		}
	}
	return nil
}

type SkillResp struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Description     *string   `json:"description,omitempty"`
}

type SkillRespArray []SkillResp