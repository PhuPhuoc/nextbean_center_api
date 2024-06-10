package model

type FilterTechnical struct {
	Id             int    `json:"id,omitempty"`
	TechnicalSkill string `json:"Technical-skill,omitempty"`
	OrderBy        string `json:"order-by,omitempty"`
}
