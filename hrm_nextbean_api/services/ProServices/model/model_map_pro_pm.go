package model

type MapProPM struct {
	ProjectId     string   `json:"project-id" validate:"required,type=string"`
	ListManagerId []string `json:"list-manager-id" validate:"required,type=string_array"`
}
