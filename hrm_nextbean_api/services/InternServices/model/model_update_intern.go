package model

type InternUpdateInfo struct {
	UserName    string `json:"user-name" validate:"required,type=string,min=4,max=20"`
	Email       string `json:"email" validate:"required,type=email"`
	StudentCode string `json:"student-code" validate:"required,type=string,min=4,max=10"`
	Avatar      string `json:"avatar" validate:"type=string"`
	Gender      string `json:"gender"  validate:"type=string"`
	DateOfBirth string `json:"date-of-birth"  validate:"type=string"`
	PhoneNumber string `json:"phone-number"  validate:"type=string"`
	Address     string `json:"address"  validate:"type=string"`
	OjtId       int    `json:"ojt-id" validate:"type=int"`
}
