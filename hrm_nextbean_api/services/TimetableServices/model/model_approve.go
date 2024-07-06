package model

type ApproveTimetable struct {
	Verified string `json:"verified" validate:"required,type=enum(denied or approved)"`
}
