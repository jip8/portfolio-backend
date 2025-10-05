package entity

import (
	"time"
	"github.com/jip/portfolio-backend"
)

type ProjectFlat struct {
	Id              *int       	`json:"id" db:"id"`
	Title          	string     	`json:"title" db:"title"`
	Description    	*string    	`json:"description" db:"description"`
	PublishedAtStr 	*string    	`json:"concluded_at" db:"concluded_at_str"`
	PublishedAt    	*time.Time 	`json:"concluded_at_time" db:"concluded_at_time"`
	Revelance      	*int       	`json:"revelance" db:"revelance"`
}

func (c *ProjectFlat) Validate() error {
	if c.Title == "" {
		return portfolio.ErrProjectTitleIsRequired
	}

	if c.PublishedAtStr != nil && *c.PublishedAtStr != "" {
		parsed, err := time.Parse("01-01-2006", *c.PublishedAtStr)
		if err != nil {
			return portfolio.ErrProjectInvalidDate
		}
		c.PublishedAt = &parsed
	}

	return nil
}

type ProjectResp struct {
	Id             int        `json:"id" db:"id"`
	Title          string     `json:"title" db:"title"`
	Description    *string    `json:"description,omitempty" db:"description"`
	PublishedAtStr *string    `json:"concluded_at" db:"concluded_at_str"`
	PublishedAt    *time.Time `json:"concluded_at_time,omitempty" db:"concluded_at_time"`
	Revelance      *int       `json:"revelance,omitempty" db:"revelance"`
}

func (c *ProjectResp) Format() error {
	if c.PublishedAt != nil {
		formatted := c.PublishedAt.Format("01-01-2006")
		c.PublishedAtStr = &formatted
		c.PublishedAt = nil
	}

	return nil
}