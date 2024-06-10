package model

type TechnicalCreationInfo struct {
	TechnicalSkill string `json:"technical-skill" validate:"required,type=string"`
}
