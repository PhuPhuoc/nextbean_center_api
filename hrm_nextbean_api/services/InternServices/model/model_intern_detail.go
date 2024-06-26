package model

type InternDetailInfo struct {
	Id            string          `json:"id"`
	UserName      string          `json:"user-name"`
	Email         string          `json:"email"`
	OjtId         int             `json:"ojt-id,omitempty"`
	StudentCode   string          `json:"student-code"`
	Avatar        string          `json:"avatar,omitempty"`
	Gender        string          `json:"gender,omitempty"`
	DateOfBirth   string          `json:"date-of-birth,omitempty"`
	PhoneNumber   string          `json:"phone-number,omitempty"`
	Address       string          `json:"address,omitempty"`
	InternSkill   []DetailSkill   `json:"detail-skill,omitempty"`
	InternProject []DetailProject `json:"detail-project,omitempty"`
}

type DetailSkill struct {
	TechnicalSkill string `json:"technical-skill"`
	SkillLevel     string `json:"skill-level"`
}

type DetailProject struct {
	ProjectName string `json:"project-name"`
	JoinAt      string `json:"join-at"`
	LeaveAt     string `json:"leave-at,omitempty"`
	Status      string `json:"status"`
}
