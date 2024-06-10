package model

type InternFilter struct {
	AccountID    string `json:"account-id,omitempty"`
	StudentCode  string `json:"student-code,omitempty"`
	OJT_Semester string `json:"ojt-semester,omitempty"`
	UserName     string `json:"user-name,omitempty"`
	Email        string `json:"email,omitempty"`
	Gender       string `json:"gender,omitempty"`
	Dob_From     string `json:"dob-from,omitempty"`
	Dob_To       string `json:"dob-to,omitempty"`
	PhoneNumber  string `json:"phone-number,omitempty"`
	Address      string `json:"address,omitempty"`
	OrderBy      string `json:"order-by,omitempty"`
}
