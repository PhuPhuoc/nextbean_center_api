package model

import "time"

type Account struct {
	Id        string    `json:"id"`
	UserName  string    `json:"user-name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created-at"`
}
