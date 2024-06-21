package model

type TaskFilter struct {
	AssgineeId   string `json:"assginee-id,omitempty"`
	AssgineeName string `json:"assginee-name,omitempty"`
	AssgineeCode string `json:"assginee-code,omitempty"`
	Status       string `json:"status,omitempty"`
	IsApproved   string `json:"is-approved,omitempty"`
	Name         string `json:"name,omitempty"`
	OrderBy      string `json:"order-by,omitempty"`
}
