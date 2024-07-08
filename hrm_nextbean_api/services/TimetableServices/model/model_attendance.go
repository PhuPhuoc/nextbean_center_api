package model

type Attendance struct {
	Clockin  string `json:"clockin" validate:"type=time"`
	Clockout string `json:"clockout" validate:"type=time"`
}
