package model

type Intern struct {
	AccountID    string `json:"account-id"`
	UserName     string `json:"user-name"`
	Email        string `json:"email"`
	StudentCode  string `json:"student-code"`
	Avatar       string `json:"avatar"`
	Gender       string `json:"gender"`
	DateOfBirth  string `json:"date-of-birth"`
	PhoneNumber  string `json:"phone-number"`
	Address      string `json:"address"`
	Ojt_semester string `json:"ojt-semester"`
}
