package model

type FilterOJT struct {
	Id         int    `json:"id,omitempty"`
	Semester   string `json:"semester,omitempty"`
	University string `json:"university,omitempty"`
	OrderBy    string `json:"order-by,omitempty"`
}
