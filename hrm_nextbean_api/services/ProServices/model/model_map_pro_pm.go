package model

type MapProPM struct {
	ListManagerId []string `json:"list-manager-id" validate:"required,type=string_array"`
}
