package entity

import (
	"time"
	"github.com/jip/portfolio-backend"
)

type ProjectFlat struct {
	Id              *int       	`json:"id" db:"id"`
	Title          	string     	`json:"title" db:"title"`
	Description    	*string    	`json:"description" db:"description"`
	PublishedAtStr 	*string    	`json:"published_at" db:"published_at_str"`
	PublishedAt    	*time.Time 	`json:"published_at_time" db:"published_at_time"`
	Revelance      	*int       	`json:"revelance" db:"revelance"`
	ThumbnailId    	*int       	`json:"thumbnail_id" db:"thumbnail_id"`
	LinksArray      LinkArray  	`json:"links" db:"-"`
	Skills          SkillArray 	`json:"skills" db:"-"`
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
	Id             int        			`json:"id" db:"id"`
	Title          string     			`json:"title" db:"title"`
	Description    *string    			`json:"description,omitempty" db:"description"`
	PublishedAtStr *string    			`json:"published_at" db:"published_at_str"`
	PublishedAt    *time.Time 			`json:"published_at_time,omitempty" db:"published_at_time"`
	Revelance      *int       			`json:"revelance,omitempty" db:"revelance"`
	ThumbnailId    *int       			`json:"thumbnail_id,omitempty" db:"thumbnail_id"`
	Thumbnail      *string    			`json:"thumbnail,omitempty" db:"thumbnail"`
	LinksRespArray LinkRespArray 		`json:"links" db:"-"`
	Attachments    *AttachmentRespArray `json:"attachments,omitempty" db:"-"`
	Skills         SkillRespArray 		`json:"skills,omitempty" db:"-"`
}

func (c *ProjectResp) Format() error {
	if c.PublishedAt != nil {
		formatted := c.PublishedAt.Format("01-01-2006")
		c.PublishedAtStr = &formatted
		c.PublishedAt = nil
	}

	return nil
}