package model

type Member struct {
	Id              string `json:"id"`
	UserName        string `json:"user-name"`
	StudentCode     string `json:"student-code"`
	Avatar          string `json:"avatar"`
	TechnicalSkills string `json:"technical_skills"`
}
