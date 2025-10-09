package entity

import (
	"github.com/jip/portfolio-backend"
)

type Contact struct {
	Id          *int       `json:"id" db:"id"`
	Link        string     `json:"link" db:"link"`
	Plataform   string     `json:"plataform" db:"plataform"`
	Description *string    `json:"description" db:"description"`
	Active      *bool      `json:"active" db:"active"`
	Revelance   *int       `json:"revelance" db:"revelance"`
}

func (c *Contact) Validate() error {
	if c.Link == "" {
		return portfolio.ErrContactLinkIsRequired
	}

	if c.Plataform == "" {
		return portfolio.ErrContactPlataformIsRequired
	}

	return nil
}

