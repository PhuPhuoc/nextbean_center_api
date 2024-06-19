package model

type Timtable struct {
	Id          string `json:"id"`
	InternName  string `json:"intern_name"`
	StudentCode string `json:"student-code"`
	OfficeTime  string `json:"office_time"`
	EstStart    string `json:"est_start"`
	EstEnd      string `json:"est_end"`
	ActStart    string `json:"act_start"`
	ActEnd      string `json:"act_end"`
	Status      string `json:"status"`
}
