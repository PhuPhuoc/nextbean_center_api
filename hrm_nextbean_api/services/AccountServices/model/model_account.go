package model

import "time"

type Account struct {
	Id          string    `json:"id"`
	UserName    string    `json:"user_name"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	Role        string    `json:"account_role"`
	CreatedTime time.Time `json:"-"`
}
