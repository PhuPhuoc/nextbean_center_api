package model

type MapInternSkill struct {
	Skills     []int    `json:"skills" validate:"required,type=int_array"`
	SkillLevel []string `json:"skill-level" validate:"required,type=string_array"`
}
