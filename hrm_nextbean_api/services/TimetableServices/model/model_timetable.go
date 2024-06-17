package model

type Timtable struct {
	Id         string `json:"id"`
	InternName string `json:"intern_name"`
	OfficeTime string `json:"office_time"`
	EstStartAt string `json:"est_start_at"`
	EstEndAt   string `json:"est_end_at"`
	ActStartAt string `json:"act_start_at,omitempty"`
	ActEndAt   string `json:"act_end_at,omitempty"`
}
