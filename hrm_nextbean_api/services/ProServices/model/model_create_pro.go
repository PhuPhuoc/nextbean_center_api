package model

type ProjectCreationInfo struct {
	Name              string `json:"name" validate:"required,type=string"`
	Description       string `json:"description" validate:"required,type=string"`
	EstStartTime      string `json:"est-start-time" validate:"required,type=date"`
	EstCompletionTime string `json:"est-completion-time" validate:"required,type=date"`
}
