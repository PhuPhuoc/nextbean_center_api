package model

type TaskUpdate struct {
	AssignedTo      string `json:"assigned-to" validate:"required,type=string"`
	IsApproved      bool   `json:"is-approved" validate:"required,type=bool"`
	Name            string `json:"name" validate:"required,type=string"`
	Description     string `json:"description" validate:"type=string"`
	EstimatedEffort int    `json:"estimated-effort" validate:"type=number"`
}
