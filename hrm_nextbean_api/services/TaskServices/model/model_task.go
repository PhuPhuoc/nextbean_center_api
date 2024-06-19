package model

type Task struct {
	Id         string `json:"id"`
	ProjectId  string `json:"project-id"`
	AssignedTo string `json:"assigned-to"`
	
}
