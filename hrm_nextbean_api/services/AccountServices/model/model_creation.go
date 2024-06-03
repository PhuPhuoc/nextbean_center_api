package model

type AccountCreationInfo struct {
	UserName string `json:"user-name" validate:"required,type=string,min=5,max=15"`
	Email    string `json:"email" validate:"required,type=email"`
	Password string `json:"password" validate:"required,type=string,min=6"`
	Role     string `json:"role" validate:"required,type=enum(admin-manager-pm-user)"`
}
