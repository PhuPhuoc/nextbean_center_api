package model

type Project struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Description string `json:"description"`
	StartDate   string `json:"start-date,omitempty"`
	Duration    string `json:"duration,omitempty"`
}
