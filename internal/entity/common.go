package entity


type List struct {
	Offset 	int          	`json:"offset"`
	Total 	int           	`json:"total"`
	Limit 	int           	`json:"limit"`
	Items 	[]interface{} 	`json:"items"`
}

type ListReq struct {
	Offset 	*int          	`json:"offset"`
	Limit 	*int           	`json:"limit"`
	Filter 	*string        	`json:"filter"`
}