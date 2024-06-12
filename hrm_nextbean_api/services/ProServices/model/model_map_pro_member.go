package model

type MapProMem struct {
	ProjectId string `json:"project-id" validate:"required,type=string"`
	MemId     string `json:"mem-id" validate:"required,type=string"`
}
