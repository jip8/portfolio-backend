package entity

import (
	"time"
	"github.com/jip/portfolio-backend"
)

type CourseFlat struct {
	Id              *int       	`json:"id" db:"id"`
	Title          	string     	`json:"title" db:"title"`
	Description    	*string    	`json:"description" db:"description"`
	ConcludedAtStr 	*string    	`json:"concluded_at" db:"concluded_at_str"`
	ConcludedAt    	*time.Time 	`json:"concluded_at_time" db:"concluded_at_time"`
	Revelance      	*int       	`json:"revelance" db:"revelance"`
}

func (c *CourseFlat) Validate() error {
	if c.Title == "" {
		return portfolio.ErrCourseTitleIsRequired
	}

	if c.ConcludedAtStr != nil && *c.ConcludedAtStr != "" {
		parsed, err := time.Parse("01-01-2006", *c.ConcludedAtStr)
		if err != nil {
			return portfolio.ErrCourseInvalidDate
		}
		c.ConcludedAt = &parsed
	}

	return nil
}

type CourseResp struct {
	Id             int        `json:"id" db:"id"`
	Title          string     `json:"title" db:"title"`
	Description    *string    `json:"description,omitempty" db:"description"`
	ConcludedAtStr *string    `json:"concluded_at" db:"concluded_at_str"`
	ConcludedAt    *time.Time `json:"concluded_at_time,omitempty" db:"concluded_at_time"`
	Revelance      *int       `json:"revelance,omitempty" db:"revelance"`
}

func (c *CourseResp) Format() error {
	if c.ConcludedAt != nil {
		formatted := c.ConcludedAt.Format("01-01-2006")
		c.ConcludedAtStr = &formatted
		c.ConcludedAt = nil
	}

	return nil
}