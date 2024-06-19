package model

type Timtable struct {
	Id          string `json:"id"`
	InternName  string `json:"intern_name"`
	StudentCode string `json:"student-code"`
	OfficeTime  string `json:"office-time"`
	EstStart    string `json:"est-start"`
	EstEnd      string `json:"est-end"`
	ActStart    string `json:"act-start"`
	ActEnd      string `json:"act-end"`
	Status      string `json:"status"`
}
