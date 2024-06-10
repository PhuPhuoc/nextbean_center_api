package model

type OJT struct {
	Id         int    `json:"id"`
	Semester   string `json:"semester"`
	University string `json:"university"`
	StartAt    string `json:"start-at"`
	EndAt      string `json:"end-at"`
}
