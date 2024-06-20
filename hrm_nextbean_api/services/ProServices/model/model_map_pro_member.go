package model

type MapProMem struct {
	MemberId string `json:"member-id" validate:"required,type=string"`
}
