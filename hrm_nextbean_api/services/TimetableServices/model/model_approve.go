package model

type ApproveTimetable struct {
	Status string `json:"status" validate:"required,type=enum(denied or approved)"`
}
