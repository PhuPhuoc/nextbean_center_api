package model

type TimeTableFilter struct {
	Id             string `json:"id,omitempty"`
	InternName     string `json:"intern-name,omitempty"`
	StudentCode    string `json:"student-code,omitempty"`
	OfficeTimeFrom string `json:"office-time-from,omitempty"`
	OfficeTimeTo   string `json:"office-time-to,omitempty"`
	Status         string `json:"status,omitempty"`
	OrderBy        string `json:"order-by,omitempty"`
}
