package entity

import (
	"time"
	"github.com/jip/portfolio-backend"
)

type ExperienceFlat struct {
	Id             	*int       	`json:"id" db:"id"`
	Title          	string     	`json:"title" db:"title"`
	Function       	*string    	`json:"function" db:"function"`
	Description    	*string    	`json:"description" db:"description"`
	InitialDateStr 	*string    	`json:"initial_date" db:"initial_date_str"`
	InitialDate    	*time.Time 	`json:"initial_date_time" db:"initial_date_time"`
	EndDateStr     	*string    	`json:"end_date" db:"end_date_str"`
	EndDate        	*time.Time 	`json:"end_date_time" db:"end_date_time"`
	Actual         	*bool      	`json:"actual" db:"actual"`
	Skills          SkillArray 	`json:"skills" db:"-"`
}


func (e *ExperienceFlat) Validate() error {
	if e.Title == "" {
		return portfolio.ErrExperienceTitleIsRequired
	}

	if e.Function == nil || *e.Function == "" {
		return portfolio.ErrExperienceFunctionIsRequired
	}

	if e.Description == nil || *e.Description == "" {
		return portfolio.ErrExperienceDescriptionIsRequired
	}

	if e.InitialDateStr != nil && *e.InitialDateStr != "" {
		parsed, err := time.Parse("01-2006", *e.InitialDateStr)
		if err != nil {
			return portfolio.ErrExperienceInvalidDate
		}
		e.InitialDate = &parsed
	}

	if e.EndDateStr != nil && *e.EndDateStr != "" {
		parsed, err := time.Parse("01-2006", *e.EndDateStr)
		if err != nil {
			return portfolio.ErrExperienceInvalidDate
		}
		e.EndDate = &parsed
	}

	if e.Actual != nil && *e.Actual && (e.EndDate != nil || e.EndDateStr != nil) {
		e.EndDate = nil
		e.EndDateStr = nil
	}

	return nil
}

type ExperienceResp struct {
	Id             int        		`json:"id" db:"id"`
	Title          string     		`json:"title" db:"title"`
	Function       *string    		`json:"function" db:"function"`
	Description    *string    		`json:"description,omitempty" db:"description"`
	InitialDateStr *string    		`json:"initial_date" db:"initial_date_str"`
	InitialDate    *time.Time 		`json:"initial_date_time,omitempty" db:"initial_date_time"`
	EndDateStr     *string    		`json:"end_date" db:"end_date_str"`
	EndDate        *time.Time 		`json:"end_date_time,omitempty" db:"end_date_time"`
	Actual         *bool      		`json:"actual,omitempty" db:"actual"`
	Skills          SkillRespArray 	`json:"skills,omitempty" db:"-"`
}

func (e *ExperienceResp) Format() error {
	if e.InitialDate != nil {
		formatted := e.InitialDate.Format("01-2006")
		e.InitialDateStr = &formatted
		e.InitialDate = nil
	}

	if e.EndDate != nil {
		formatted := e.EndDate.Format("01-2006")
		e.EndDateStr = &formatted
		e.EndDate = nil
	}

	return nil
}
