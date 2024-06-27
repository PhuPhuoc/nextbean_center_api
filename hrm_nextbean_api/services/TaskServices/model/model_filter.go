package model

type TaskFilter struct {
	ProjectId    string `json:"-"`
	AssigneeId   string `json:"assignee-id,omitempty"`
	AssigneeName string `json:"assignee-name,omitempty"`
	AssigneeCode string `json:"assignee-code,omitempty"`
	Status       string `json:"status,omitempty"`
	IsApproved   string `json:"is-approved,omitempty"`
	Name         string `json:"name,omitempty"`
	OrderBy      string `json:"order-by,omitempty"`
}
