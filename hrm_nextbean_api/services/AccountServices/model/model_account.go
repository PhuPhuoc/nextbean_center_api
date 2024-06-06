package model

type Account struct {
	Id        string `json:"id"`
	UserName  string `json:"user-name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Role      string `json:"role"`
	CreatedAt string `json:"created-at,omitempty"`
}
