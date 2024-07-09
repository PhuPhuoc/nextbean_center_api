package model

type StatusAbsent struct {
	Status string `json:"status" validate:"required,type=enum(absent or remove)"`
}
