package model

type TaskUpdate struct {
	AssignedTo      string `json:"assigned-to" validate:"required,type=string"`
	IsApproved      bool   `json:"is-approved" validate:"required,type=bool"`
	Status          string `json:"status" validate:"required,type=enum(inprogress or done)"`
	Name            string `json:"name" validate:"required,type=string"`
	Description     string `json:"description" validate:"type=string"`
	EstimatedEffort string `json:"estimated-effort" validate:"type=string"`
}
