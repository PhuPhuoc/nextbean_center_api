package model

type AccountFilter struct {
	Id            string `json:"id,omitempty"`
	UserName      string `json:"user-name,omitempty"`
	Email         string `json:"email,omitempty"`
	Role          string `json:"role,omitempty"`
	CreatedAtFrom string `json:"created-at-from,omitempty"`
	CreatedAtTo   string `json:"created-at-to,omitempty"`
	OrderBy       string `json:"order-by,omitempty"`
}
