package entity

type ExperienceFlat struct {
	Id          *int    `json:"id"`
	Title       string `json:"title"`
	Description *string `json:"description"`
}

type ExperienceResp struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}