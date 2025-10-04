package entity

type List[T any] struct {
	Offset int `json:"offset"`
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Items  []T `json:"items"`
}

type ListReq struct {
	Offset *int    `json:"offset"`
	Limit  *int    `json:"limit"`
	Filter *string `json:"filter"`
}