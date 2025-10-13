package entity

import (
	"github.com/jip/portfolio-backend"
)

type LinkFlat struct {
	Id              *int       `json:"id" db:"id"`
	ParentId        *int       `json:"parent_id" db:"parent_id"`
	Module          *string    `json:"module" db:"module"`
	Title           *string    `json:"title" db:"title"`
	Link            string     `json:"link" db:"link"`
	Revelance       *int       `json:"revelance" db:"revelance"`
	Description     *string    `json:"description" db:"description"`
	UpdatedAt       *string    `json:"updated_at" db:"updated_at"`
}

type LinkArray []LinkFlat

func (l LinkArray) Validate() error {
	for _, link := range l {
		if link.Id == nil {
			return portfolio.ErrLinkIdIsRequired
		}

		if link.Link == "" {
			return portfolio.ErrLinkLinkIsRequired
		}
	}
	return nil
}

type LinkResp struct {
	Id              int       `json:"id"`
	Title           *string   `json:"title"`
	Link            string    `json:"link"`
	Description     *string   `json:"description,omitempty"`
}

type LinkRespArray []LinkResp