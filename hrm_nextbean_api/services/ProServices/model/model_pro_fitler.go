package model

type ProjectFilter struct {
	Name          string `json:"name,omitempty"`
	Status        string `json:"status,omitempty"`
	StartDateFrom string `json:"start-date-from,omitempty"`
	StarttDateTo  string `json:"start-date-to,omitempty"`
	OrderBy       string `json:"order-by,omitempty"`
	Role          string `json:"-"`
	AccId         string `json:"-"`
}
