package model

type UpdateProjectInfo struct {
	Name              string `json:"name" validate:"required,type=string"`
	Status            string `json:"status" validate:"required,type=enum(not_started or in_progress or completed or cancel)"`
	Description       string `json:"description" validate:"required,type=string"`
	EstStartTime      string `json:"est-start-time" validate:"required,type=date"`
	EstCompletionTime string `json:"est-completion-time" validate:"required,type=date"`
}
