package content

type Posts struct {
	Meta  Meta   `json:"meta"`
	Posts []Post `json:"posts"`
}
