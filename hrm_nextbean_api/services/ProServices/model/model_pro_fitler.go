package model

type ProjectFilter struct {
	Name                  string `json:"name,omitempty"`
	Status                string `json:"status,omitempty"`
	EstStartTimeFrom      string `json:"est-start-time-from,omitempty"`
	EstStartTimeTo        string `json:"est-start-time-to,omitempty"`
	EstCompletionTimeFrom string `json:"est-completion-time-from,omitempty"`
	EstCompletionTimeTo   string `json:"est-completion-time-to,omitempty"`
	OrderBy               string `json:"order-by,omitempty"`
	Role                  string `json:"-"`
	AccId                 string `json:"-"`
}
