package model

type MapInternSkill struct {
	InternId   string   `json:"intern-id" validate:"required,type=string"`
	Skills     []int    `json:"skills" validate:"required,type=int_array"`
	SkillLevel []string `json:"skill-level" validate:"required,type=string_array"`
}
