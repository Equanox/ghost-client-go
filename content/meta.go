package content

type Meta struct {
	Pagination struct {
		Limit int         `json:"limit"`
		Next  interface{} `json:"next"`
		Page  int         `json:"page"`
		Pages int         `json:"pages"`
		Prev  interface{} `json:"prev"`
		Total int         `json:"total"`
	} `json:"pagination"`
}
