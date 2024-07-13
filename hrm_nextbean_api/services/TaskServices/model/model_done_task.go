package model

type DoneTask struct {
	ActualEffort int `json:"actual-effort" validate:"type=number"`
}
