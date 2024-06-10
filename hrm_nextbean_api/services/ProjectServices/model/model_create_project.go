package model

type ProjectCreationInfo struct {
	Name        string `json:"name" validate:"required,type=string"`
	Description string `json:"description" validate:"required,type=string"`
	Duration    string `json:"duration" validate:"required,type=string"`
	StartAt     string `json:"start-at" validate:"required,type=date"`
}
