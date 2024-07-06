package model

type Task struct {
	Id              string `json:"id"`
	ProjectId       string `json:"project-id"`
	ProjectName     string `json:"project-name"`
	AssignedTo      string `json:"assigned-to"`
	AssignedName    string `json:"assigned-name"`
	AssignedCode    string `json:"assigned-code"`
	IsApproved      string `json:"is-approved"`
	Status          string `json:"status"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	EstimatedEffort string `json:"estimated-effort"`
	ActualEffort    string `json:"actual-effort"`
}
