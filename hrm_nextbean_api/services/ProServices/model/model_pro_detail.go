package model

type ProjectDetail struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	Description        string `json:"description"`
	EstStartTime       string `json:"est-start-time,omitempty"`
	EstCompletionTime  string `json:"est-completion-time,omitempty"`
	TotalMember        int    `json:"total-member"`
	TotalTask          int    `json:"total-task"`
	TotalTaskCompleted int    `json:"total-task-conpleted"`
}
