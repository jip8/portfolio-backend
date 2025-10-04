package entity

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