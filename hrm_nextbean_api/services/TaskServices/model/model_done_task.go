package model

type DoneTask struct {
	ActualEffort string `json:"actual-effort" validate:"type=string"`
}
