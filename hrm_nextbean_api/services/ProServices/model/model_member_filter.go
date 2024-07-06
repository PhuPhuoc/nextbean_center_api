package model

type MemberFilter struct {
	UserName    string `json:"user-name,omitempty"`
	StudentCode string `json:"student-code,omitempty"`
	Semester    string `json:"semester,omitempty"`
	University  string `json:"university,omitempty"`
	Status      string `json:"status,omitempty"`
}
