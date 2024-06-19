package model

type TaskCreation struct {
	AssignedTo      string `json:"assigned-to" validate:"required,type=string"`
	Name            string `json:"name" validate:"required,type=string"`
	Description     string `json:"description" validate:"type=string"`
	EstimatedEffort string `json:"estimated-effort" validate:"type=string"`
}
