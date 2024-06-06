package model

type UpdateAccountInfo struct {
	Id       string `json:"id" validate:"required,type=string"`
	UserName string `json:"user-name" validate:"required,type=string,min=4,max=20"`
	Email    string `json:"email" validate:"required,type=email"`
	Role     string `json:"role" validate:"required,type=enum(admin or manager or pm or user)"`
}
