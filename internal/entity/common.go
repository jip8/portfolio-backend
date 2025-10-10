package entity

type List[T any] struct {
	Offset int `json:"offset"`
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Items  []T `json:"items"`
}

type ListReq struct {
	Order  string `json:"order"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Filter string `json:"filter"`
}

func (r *ListReq) Process() {
	if r.Limit == 0 {
		r.Limit = 100
	}

	// TODO: validate order
	if r.Order == "" {
		r.Order = "id ASC"
	}
}