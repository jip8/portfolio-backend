package entity

import (
	"time"

	"github.com/jip/portfolio-backend"
)

type ArticleFlat struct {
	Id              *int       	`json:"id" db:"id"`
	Title          	string     	`json:"title" db:"title"`
	Description    	*string    	`json:"description" db:"description"`
	Local           *string    	`json:"local" db:"local"`
	PublishedAtStr 	*string    	`json:"published_at" db:"published_at_str"`
	PublishedAt    	*time.Time 	`json:"published_at_time" db:"published_at_time"`
	Revelance      	*int       	`json:"revelance" db:"revelance"`
	LinksArray      LinkArray  	`json:"links" db:"-"`
}

func (c *ArticleFlat) Validate() error {
	if c.Title == "" {
		return portfolio.ErrArticleTitleIsRequired
	}

	if c.PublishedAtStr != nil && *c.PublishedAtStr != "" {
		parsed, err := time.Parse("01-01-2006", *c.PublishedAtStr)
		if err != nil {
			return portfolio.ErrArticleInvalidDate
		}
		c.PublishedAt = &parsed
	}

	return nil
}

type ArticleResp struct {
	Id              int        		`json:"id" db:"id"`
	Title          	string     		`json:"title" db:"title"`
	Description    	*string    		`json:"description" db:"description"`
	Local           *string    		`json:"local" db:"local"`
	PublishedAtStr 	*string    		`json:"published_at,omitempty" db:"published_at_str"`
	PublishedAt    	*time.Time 		`json:"published_at_time,omitempty" db:"published_at_time"`
	Revelance      	*int       		`json:"revelance,omitempty" db:"revelance"`
	LinksRespArray  LinkRespArray 	`json:"links,omitempty" db:"-"`
}

func (c *ArticleResp) Format() error {
	if c.PublishedAt != nil {
		formatted := c.PublishedAt.Format("01-01-2006")
		c.PublishedAtStr = &formatted
		c.PublishedAt = nil
	}

	return nil
}