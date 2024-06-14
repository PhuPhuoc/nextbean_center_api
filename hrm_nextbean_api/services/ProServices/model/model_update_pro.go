package model

type UpdateProjectInfo struct {
	Name        string `json:"name" validate:"required,type=string"`
	Status      string `json:"status" validate:"required,type=enum(not_start or doing or done or cancel)"`
	Description string `json:"description" validate:"required,type=string"`
	Duration    string `json:"duration" validate:"required,type=string"`
	StartDate   string `json:"start-date" validate:"required,type=date"`
}
