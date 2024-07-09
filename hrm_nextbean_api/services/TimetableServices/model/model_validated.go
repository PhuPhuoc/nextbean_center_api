package model

type AttendanceValidated struct {
	ValidateField string `json:"validate-field" validate:"required,type=enum(clockin or clockout)"`
}
