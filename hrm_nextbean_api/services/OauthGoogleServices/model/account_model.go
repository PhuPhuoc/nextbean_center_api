package model

type AccountLoginGG struct {
	Id        string `json:"id"`
	UserName  string `json:"user-name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created-at,omitempty"`
}
