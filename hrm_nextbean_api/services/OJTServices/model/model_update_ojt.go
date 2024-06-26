package model

type UpdateOJTInfo struct {
	Semester   string `json:"semester" validate:"required,type=string"`
	University string `json:"university" validate:"required,type=string"`
	StartAt    string `json:"start-at" validate:"required,type=date"`
	EndAt      string `json:"end-at" validate:"required,type=date"`
}
