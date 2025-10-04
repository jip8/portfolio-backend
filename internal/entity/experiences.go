package entity

import "time"

type Experience struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Function    string    `db:"function" json:"function"`
	Description string    `db:"description" json:"description"`
	InitialDate time.Time `db:"initial_date" json:"initialDate"`
	EndDate     time.Time `db:"end_date" json:"endDate"`
}

type ExperienceFlat struct {
	Id          *int    `json:"id"`
	Title       string 	`json:"title"`
	Function    *string `json:"function"`
	Description *string `json:"description"`
	InitialDate *string `json:"initialDate"`
	EndDate     *string `json:"endDate"`
}

type ExperienceResp struct {
	Id          string 	`json:"id"`
	Title       string 	`json:"title"`
	Function    *string `json:"function"`
	Description *string `json:"description"`
	InitialDate *string `json:"initialDate"`
	EndDate     *string `json:"endDate"`
}