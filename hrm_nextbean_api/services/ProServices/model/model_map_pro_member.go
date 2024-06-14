package model

type MapProMem struct {
	MemId string `json:"mem-id" validate:"required,type=string"`
}
