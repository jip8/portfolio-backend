package entity

type ExperienceFlat struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ExperienceResp struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}